import * as React from "react";
import {useEffect, useRef, useState} from "react";
import {Button} from "primereact/button";
import {useRouter} from "next/router";
import useWebSocket from 'react-use-websocket';
import axios from "axios";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faArrowRightFromBracket, faCaretDown, faPhone, faUserSecret} from "@fortawesome/free-solid-svg-icons";

import {OverlayPanel} from "primereact/overlaypanel";
import {Menu} from "primereact/menu";
import {InputTextarea} from "primereact/inputtextarea";
import Chat from "./chat";
import Moment from "react-moment";
import WebRTC from "../../components/webrtc/WebRTC";
import {FeedItem} from "../../model/feed-item";
import {toast, ToastContainer} from "react-toastify";
import CallProgress from '../../components/webrtc/CallProgress';
import SuggestionList, {Suggestion} from "../../components/SuggestionList";

import {gql} from "graphql-request";
import {useGraphQLClient} from "../../utils/graphQLClient";
import {Checkbox} from "primereact/checkbox";

interface FeedProps {
  feedId: string | undefined;
  userLoggedInEmail: string;
  logoutUrl: string | undefined;
}

export const Feed = (props: FeedProps) => {
  const router = useRouter()
  const client = useGraphQLClient();

  const [feeds, setFeeds] = useState([] as FeedItem[]);
  const [selectedFeed, setSelectedFeed] = useState(props.feedId);
  const [dialedNumber, setDialedNumber] = useState('');
  const [checked, setChecked] = useState(false);

  const handleCheckboxChange = (event: { checked: boolean }) => {
    setChecked(event.checked);
    loadFeed(event.checked)
  };

  const {lastMessage} = useWebSocket(`${process.env.NEXT_PUBLIC_WEBSOCKET_PATH}`, {
    onOpen: () => console.log('Websocket opened'),
    //Will attempt to reconnect on all close events, such as server shutting down
    shouldReconnect: (closeEvent) => true,
  });

  useEffect(() => {
    loadFeed(false);
  }, []);

  useEffect(() => {
    if (lastMessage && Object.keys(lastMessage).length !== 0) {
      loadFeed(false);
    }

  }, [lastMessage]);

  const loadFeed = function (onlyContacts: boolean) {
    console.log("Reloading feed!");
    axios.get(`/oasis-api/feed?onlyContacts=${onlyContacts}`)
      .then(res => {
        setFeeds(res.data?.feedItems ?? []);
        if (!selectedFeed && res.data && res.data.feedItems && res.data.feedItems[0]) {
          setSelectedFeed(res.data.feedItems[0].id);
          router.push(`/feed?id=${res.data.feedItems[0].id}`, undefined, {shallow: true});
        }
      });
  }

  const [referStatus, setReferStatus] = useState(props.feedId);

  const userSettingsContainerRef = useRef<OverlayPanel>(null);

  let userItems = [
    {
      label: 'My profile',
      icon: <FontAwesomeIcon icon={faUserSecret} className="mr-2"/>,
      command: () => {
        router.push('/');
      }
    },
    {
      label: 'Logout',
      icon: <FontAwesomeIcon icon={faArrowRightFromBracket} className="mr-2"/>,
      command: () => {
        window.location.href = props.logoutUrl as string;
      }
    }
  ];

  // region WebRTC
  const webrtc: React.RefObject<WebRTC> = useRef<WebRTC>(null);

  useEffect(() => {

    const refreshCredentials = () => {
      axios.get(`/oasis-api/call_credentials?service=sip`)
        .then(res => {
          console.error("Got a key: " + JSON.stringify(res.data));

          webrtc.current?.setCredentials(res.data.username, res.data.password,
            () => {
              if (!webrtc.current?._ua) {
                webrtc.current?.startUA()
              }
            });
          setTimeout(() => {
            refreshCredentials()
          }, (res.data.ttl * 3000) / 4);
        });
    }
    refreshCredentials();
  }, []);

  const [callFrom, setCallFrom] = useState('');
  const [inCall, setInCall] = useState(false);

  interface FeedInitatior {
    phoneNumber?: string
    email?: string
  }

  const buildFeedInitiator = (dest: string): FeedInitatior => {
    const feedInitiator: FeedInitatior = {};
    if (dest.includes("@")) {
      feedInitiator.email = dest;
    } else {
      feedInitiator.phoneNumber = dest;
    }
    return feedInitiator;
  }

  const handleCall = (feedInitiator: FeedInitatior) => {
    let user = '';
    if (feedInitiator.phoneNumber) {
      user = feedInitiator.phoneNumber + "@oasis.openline.ai";
    } else if (feedInitiator.email) {
      user = feedInitiator.email;
      const regex = /.*<(.*)>/;
      const matches = user.match(regex);
      if (matches) {
        user = matches[1];
      }
    }
    webrtc.current?.makeCall("sip:" + user);
  }

  function setTopbarColor(newColor: any) {
    document.documentElement.style.setProperty('--topbar-background', newColor);
  }

  useEffect(() => {
    if (inCall) {
      setTopbarColor('#FFCCCB');
    } else {
      setTopbarColor('#FFFFFF');
    }
  }, [inCall]);

  interface ContactResponse {
    contacts: {
      content: {
        id: string;
        firstName: string;
        lastName: string;
        phoneNumbers: { e164: string }[];
      }[]
    }

  }

  const getContactSuggestions = (filter: string, callback: Function) => {
    const query = gql`query  getContacts($value: Any!) {
            contacts( where: {OR: [{filter: {property: "FIRST_NAME", value: $value, operation: CONTAINS }}, {filter: {property: "LAST_NAME", value: $value, operation: CONTAINS }}]})
            {
                content{id, firstName, lastName, phoneNumbers{e164}}
            }
        }`

    client.request(query, {value: filter}).then((response: ContactResponse) => {
      var suggestions: Suggestion[] = [];
      if (response.contacts && response.contacts.content) {
        for (const contact of response.contacts.content) {
          if (contact.phoneNumbers && contact.phoneNumbers.length > 0) {
            var sugestion = {
              id: contact.id,
              display: contact.firstName + " " + contact.lastName,
              value: contact.phoneNumbers[0].e164
            }
            suggestions.push(sugestion);
          }
        }
      }
      callback(suggestions);

    }).catch(reason => {
      toast.error("There was a problem on our side and we are doing our best to solve it!");
    });
  }

  return (
    <div className="flex w-full h-full">
      <ToastContainer position="top-center"
                      autoClose={3000}
                      closeOnClick={true}
                      hideProgressBar={true}
                      theme="colored"/>
      {
        process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL &&
        <WebRTC
          ref={webrtc}
          websocket={`${process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL}`}
          from={"sip:" + props.userLoggedInEmail}
          notifyCallFrom={(callFrom: string) => setCallFrom(callFrom)}
          updateCallState={(state: boolean) => setInCall(state)}
          updateReferStatus={(status: string) => setReferStatus(status)}
          autoStart={false}
        />
      }

      <div className="flex flex-column flex-grow-0 h-full overflow-hidden"
           style={{width: '350px', background: 'white', borderRight: '1px rgb(235, 235, 235) solid'}}>
        <div>
          <div className='flex p-3 w-full'>
            <InputTextarea className="w-full"
                           value={dialedNumber}
                           onChange={(e) => setDialedNumber(e.target.value)}
                           autoResize
                           rows={1}
                           placeholder="Call"
                           onKeyPress={(e) => {
                             if (e.shiftKey && e.key === "Enter") {
                               return true
                             }
                             if (e.key === "Enter") {
                               e.preventDefault();
                             }
                           }}
                           style={{
                             borderColor: "black", //Do not set as none!! It breaks InputTextarea autoResize
                             boxShadow: "none"
                           }}
            />
            <div className="flex p-1">
              <Checkbox
                inputId="checkboxId"
                onChange={handleCheckboxChange}
                checked={checked}
              />
            </div>
            {dialedNumber.length > 0 && !inCall &&
              <Button
                onClick={() => handleCall(buildFeedInitiator(dialedNumber))}
                tooltip={
                  `Call (${dialedNumber})`
                }
                tooltipOptions={{position: 'top', showDelay: 200, hideDelay: 200}}
                className='p-button-text mx-2 p-2'
                style={{
                  border: 'solid 1px #E8E8E8',
                  borderRadius: '6px'
                }}>
                <FontAwesomeIcon icon={faPhone} style={{fontSize: '20px'}}/>
              </Button>
            }
          </div>

          <div className='flex p-3'>
            <SuggestionList currentValue={dialedNumber}
                            getSuggestions={getContactSuggestions}
                            setCurrentValue={setDialedNumber}></SuggestionList>
          </div>
        </div>


        <div className='flex flex-column pl-3 pr-3 mb-3 overflow-x-hidden overflow-y-auto'>
          {
            feeds.map((f: FeedItem) => {
              let className = 'flex w-full align-content-center align-items-center p-3 mb-2 contact-hover';
              if (selectedFeed === f.id) {
                className += ' selected'
              }

              var t = new Date(1970, 0, 1);
              t.setSeconds(f.lastTimestamp.seconds);

              return <div key={f.id} className={className} onClick={() => {
                setSelectedFeed(f.id);
                //change the URL to allow a bookmark
                router.push(`/feed?id=${f.id}`, undefined, {shallow: true});
              }
              }>
                <div className='flex flex-column flex-grow-1 mr-3' style={{minWidth: '0'}}>
                  <div className='mb-2'>
                    {
                      f.initiatorFirstName &&
                      f.initiatorFirstName + ' ' + f.initiatorLastName}
                    {
                      !f.initiatorFirstName &&
                      f.initiatorUsername
                    }
                  </div>
                  <div className='text-500' style={{
                    fontSize: '12px',
                    textOverflow: 'ellipsis',
                    whiteSpace: "nowrap",
                    overflow: "hidden"
                  }}>
                    {f.lastContentPreview}
                  </div>
                </div>

                <div className='flex flex-column text-right'>
                  <Moment className="text-sm text-gray-600" date={t}
                          format={'MMM D, YYYY'}></Moment>
                  <Moment className="text-sm text-gray-600" date={t}
                          format={'HH:mma'}></Moment>
                </div>
              </div>
            })
          }
        </div>

      </div>

      <div className='flex h-full w-full flex-column'>
        <div className='openline-top-bar'>
          <div className="flex align-items-center justify-content-end">

            <CallProgress inCall={inCall} webrtc={webrtc} callFrom={callFrom}
                          referStatus={referStatus as string}
                          getContactSuggestions={getContactSuggestions}></CallProgress>
            <Button className="flex-none px-3 m-3"
                    onClick={(e: any) => userSettingsContainerRef?.current?.toggle(e)}>
              <FontAwesomeIcon icon={faUserSecret} className="mr-2"/>
              <span className='flex-grow-1'>{props.userLoggedInEmail}</span>
              <FontAwesomeIcon icon={faCaretDown} className="ml-2"/>
            </Button>

            <OverlayPanel ref={userSettingsContainerRef} dismissable>
              <Menu model={userItems} style={{border: 'none'}}/>
            </OverlayPanel>

          </div>
        </div>
        {
          selectedFeed &&
          <Chat
            feedId={selectedFeed}
            inCall={inCall}
            userLoggedInEmail={props.userLoggedInEmail}
            handleCall={(feedInitiator: any) => handleCall(feedInitiator)}
          />
        }
      </div>


    </div>
  );
}

export default Feed;