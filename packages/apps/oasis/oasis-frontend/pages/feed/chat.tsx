import * as React from "react";
import {useEffect, useState} from "react";
import {Button} from "primereact/button";
import {SplitButton} from 'primereact/splitbutton';
import {faMinus, faPaperclip, faPhone, faPlus, faSmile} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {InputTextarea} from "primereact/inputtextarea";
import axios from "axios";
import useWebSocket from "react-use-websocket";
import {gql} from "graphql-request";
import {ProgressSpinner} from "primereact/progressspinner";
import {Tooltip} from 'primereact/tooltip';
import Moment from "react-moment";
import {FeedItem} from "../../model/feed-item";
import {toast} from "react-toastify";
import {ConversationItem, FeedPostRequest} from "../../model/conversation-item";
import {useGraphQLClient} from "../../utils/graphQLClient";
import sanitizeHtml from "sanitize-html";

interface ChatProps {
  feedId: string;
  inCall: boolean;
  userLoggedInEmail: string;

  handleCall(feedInitiator: any): void;
}

interface Participant {
  email: string;
}

export const Chat = (props: ChatProps) => {
  const client = useGraphQLClient();

  const {lastMessage} = useWebSocket(`${process.env.NEXT_PUBLIC_WEBSOCKET_PATH}/${props.feedId}`, {
    onOpen: () => console.log('Websocket opened'),
    //Will attempt to reconnect on all close events, such as server shutting down
    shouldReconnect: (closeEvent) => true,
  })

  const [feedInitiator, setFeedInitiator] = useState({
    loaded: false,
    email: '',
    firstName: '',
    lastName: '',
    phoneNumber: ''
  });

  function decodeChannel(channel: number) {
    switch (channel) {
      case 0:
        return "Web chat";
      case 1:
        return "Email";
      case 2:
        return "WhatsApp";
      case 3:
        return "Facebook";
      case 4:
        return "Twitter";
      case 5:
        return "Phone call";
    }
    return "";
  }

  function callingAllowed() {
    return process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL && (feedInitiator.phoneNumber || feedInitiator.email == "echo@oasis.openline.ai");
  }

  const [currentChannel, setCurrentChannel] = useState('CHAT');
  const [currentText, setCurrentText] = useState('');
  const [addParticipantText, setAddParticipantText] = useState('');

  const [sendButtonDisabled, setSendButtonDisabled] = useState(false);
  const [messages, setMessages] = useState([] as ConversationItem[]);
  const [participants, setParticipants] = useState([] as Participant[]);


  const [loadingMessages, setLoadingMessages] = useState(false)

  useEffect(() => {
    if (props.feedId) {
      setLoadingMessages(true);
      setCurrentText('');

      axios.get(`/oasis-api/feed/${props.feedId}`)
        .then(res => {
          const feedItem = res.data as FeedItem;

          if (feedItem.initiatorType === 'CONTACT') {

            const queryMail = gql`query GetContactDetails($email: String!) {
                        contact_ByEmail(email: $email) {
                            id
                            firstName
                            lastName
                            emails {
                                email
                            }
                            phoneNumbers {
                                e164
                            }
                        }
                    }`
            const queryPhone = gql`query GetContactDetails($phone: String!) {
                      contact_ByPhone(e164: $phone) {
                          id
                          firstName
                          lastName
                          emails {
                              email
                          }
                          phoneNumbers {
                              e164
                          }
                      }
                  }`
            
            
            if (feedItem.initiatorUsername.type === 0) {     
              client.request(queryMail, {email: feedItem.initiatorUsername.identifier}).then((response: any) => {
                if (response.contact_ByEmail) {
                  setFeedInitiator({
                    loaded: true,
                    firstName: response.contact_ByEmail.firstName,
                    lastName: response.contact_ByEmail.lastName,
                    email: response.contact_ByEmail.emails[0]?.email ?? undefined,
                    phoneNumber: response.contact_ByEmail.phoneNumbers[0]?.e164 ?? undefined
                  });
                } else {
                  //todo log on backend
                  toast.error("There was a problem on our side and we are doing our best to solve it!");
                }
              }).catch(reason => {
                //todo log on backend
                toast.error("There was a problem on our side and we are doing our best to solve it!");
              });
          } else if (feedItem.initiatorUsername.type === 1) {
              client.request(queryPhone, {phone: feedItem.initiatorUsername.identifier}).then((response: any) => {
                if (response.contact_ByPhone) {
                  setFeedInitiator({
                    loaded: true,
                    firstName: response.contact_ByPhone.firstName,
                    lastName: response.contact_ByPhone.lastName,
                    email: response.contact_ByPhone.emails[0]?.email ?? undefined,
                    phoneNumber: response.contact_ByPhone.phoneNumbers[0]?.e164 ?? undefined
                  });
                } else {
                  //todo log on backend
                  toast.error("There was a problem on our side and we are doing our best to solve it!");
                }
              }).catch(reason => {
                //todo log on backend
                toast.error("There was a problem on our side and we are doing our best to solve it!");
              });
            }
            //TODO move initiator in index
          } else if (feedItem.initiatorType === 'USER') {

            const query = gql`query GetUserDetails($email: String!)  {
              user_ByEmail(email: $email) {
                            id
                            firstName
                            lastName
                            emails {
                              email
                          }
                        }
                    }`

            client.request(query, {email: feedItem.initiatorUsername.identifier}).then((response: any) => {
              if (response.user_ByEmail) {
                setFeedInitiator({
                  loaded: true,
                  firstName: response.user_ByEmail.firstName,
                  lastName: response.user_ByEmail.lastName,
                  email: response.user_ByEmail.emails[0]?.email ?? "undefined",
                  phoneNumber:  '' //TODO user doesn't have phone in backend
                });
              } else {
                //TODO log on backend
                toast.error("There was a problem on our side and we are doing our best to solve it!");
              }
            }).catch(reason => {
              //TODO log on backend
              toast.error("There was a problem on our side and we are doing our best to solve it! " + reason);
            });

          }

        }).catch((reason: any) => {
        //todo log on backend
        toast.error("There was a problem on our side and we are doing our best to solve it!");
      });

      axios.get(`/oasis-api/feed/${props.feedId}/item`)
        .then(res => {
          setMessages(res.data ?? []);
        }).catch((reason: any) => {
        //todo log on backend
        toast.error("There was a problem on our side and we are doing our best to solve it!");
      });
      axios.get(`/oasis-api/feed/${props.feedId}/participants`)
        .then(res => {
          if (res.data && res.data.participants) {
            let list: Participant[] = [];
            for (let participant of res.data.participants) {
              if (participant.email !== props.userLoggedInEmail) {
                list.push({email: participant.email});
              }
            }
            setParticipants(list);
          }
        }).catch((reason: any) => {
        //todo log on backend
        toast.error("There was a problem on our side and we are doing our best to solve it!");
      });
    }
  }, [props.feedId]);

  //when a new message appears, scroll to the end of container
  useEffect(() => {
    if (messages && feedInitiator.loaded) {
      setLoadingMessages(false);
    }
  }, [messages, feedInitiator]);

  useEffect(() => {
    if (!loadingMessages) {
      const element = document.getElementById('chatWindowToScroll')
      if (element) {
        element.scrollIntoView({block: 'end', inline: 'nearest'}) // TODO: add separate behaviour for new messages that has a smooth scroll (behavior: 'smooth')
      }
    }
  }, [loadingMessages])

  useEffect(() => {
    setSendButtonDisabled(currentText === '')
  }, [currentText]);

  useEffect(() => {
    if (lastMessage && Object.keys(lastMessage).length !== 0 && lastMessage.data.length > 0) {
      handleWebsocketMessage(JSON.parse(lastMessage.data));
    }
  }, [lastMessage]);

  const handleSendMessage = () => {
    if (!currentText) return;
    let message: FeedPostRequest = {
      channel: currentChannel,
      username: props.userLoggedInEmail,
      message: currentText,
      direction: 'OUTBOUND',
      destination: participants.map((participant: Participant) => participant.email)
    }
    if (messages.length > 0) {
      message.replyTo = messages[messages.length - 1].messageId.conversationEventId;
    }
    axios.post(`/oasis-api/feed/${props.feedId}/item`, message).then(res => {
      console.log(res)
      if (res.data) {
        setMessages((messageList: any) => [res.data, ...messageList]);
        setCurrentText('');
      }
    }).catch(reason => {
      //todo log on backend
      toast.error("There was a problem on our side and we are doing our best to solve it!");
    });
  };

  const handleWebsocketMessage = function (msg: any) {
    let newMsg: ConversationItem = {
      content: msg.content,
      senderUsername: msg.SenderUserName.identifier,
      type: msg.Type,
      time: msg.time,
      messageId: msg.messageId,
      direction: msg.direction == "OUTBOUND" ? 1 : 0,
      subtype: 0,
      senderType: 0,
      senderId: ""
    };

    setMessages((messageList: any) => [newMsg, ...messageList]);
  }

  const showParticipants = () => {
    return participants.map((participant: Participant, index: any) => {
      return <><span style={{
        border: 'solid 1px #E8E8E8',
        borderRadius: '8px',
        boxShadow: '0px 0px 40px rgba(0, 0, 0, 0.05)',
        padding: '5px',
        background: '#E8E8E8',
        verticalAlign: 'sub',
      }}>
        <Button className='p-button-text'>
          <FontAwesomeIcon icon={faMinus} style={{fontSize: '20px'}}
                           onClick={() => setParticipants(participants.filter(e => e.email !== participant.email))}/>
        </Button>
        <span>
          {participant.email}
        </span>
      </span>
      </>
    })
  }

  const addParticipantBox = () => {
    return <>
      <span style={{
        border: 'solid 1px #E8E8E8',
        borderRadius: '8px',
        boxShadow: '0px 0px 40px rgba(0, 0, 0, 0.05)',
        padding: '5px',
        background: '#E8E8E8',
        verticalAlign: 'sub',
      }}>
        <Button className='p-button-text'>
          <FontAwesomeIcon icon={faPlus}
                           style={{fontSize: '20px'}}
                           onClick={() => {
                             setParticipants([...participants, {email: addParticipantText}]);
                             setAddParticipantText("")
                           }}/>
        </Button>
        <input type="text"
               name="newParticipant"
               value={addParticipantText}
               onChange={(e) => setAddParticipantText(e.target.value)}
        />
        </span>
    </>
  }
  const sendButtonOptions = [
    {
      label: 'Web chat',
      value: 'CHAT',
      command: (e: any) => {
        setCurrentChannel(e.item.value)
      }
    },
    {
      label: 'Email',
      value: 'EMAIL',
      command: (e: any) => {
        setCurrentChannel(e.item.value)
      }
    },
  ]

  return (
    <div className='flex flex-column h-full w-full overflow-hidden'>
      <div className="flex-grow-1 w-full overflow-x-hidden overflow-y-auto p-5 pb-0">
        {
          loadingMessages &&
          <div className="flex w-full h-full align-content-center align-items-center">
            <ProgressSpinner/>
          </div>
        }

        <div className="flex flex-column">
          {
            !loadingMessages &&
            messages.reverse().map((msg: ConversationItem, index: any) => {
              let lines = msg.content.split('\n');

              let filtered: string[] = lines.filter(function (line: string) {
                return line.indexOf('>') != 0;
              });
              msg.content = filtered.join('\n').trim();

              var t = new Date(1970, 0, 1);
              t.setSeconds(msg.time.seconds);

              return <div key={msg.messageId.conversationEventId}
                          className='flex flex-column mb-3'>
                {
                  msg.direction == 0 &&
                  <>
                    {
                      (index == 0 || (index > 0 && messages[index - 1].direction !== messages[index].direction)) &&
                      <div className="mb-1 pl-3">
                        {
                          feedInitiator.firstName && feedInitiator.lastName &&
                          <>{feedInitiator.firstName} {feedInitiator.lastName}</>
                        }
                        {
                          !feedInitiator.firstName && !feedInitiator.lastName &&
                          <>{feedInitiator.email}</>
                        }
                      </div>
                    }

                    <div className="flex">
                      <div className="flex flex-column flex-grow-0 p-3" style={{
                        background: 'white',
                        borderRadius: '5px',
                        boxShadow: '0 2px 1px -1px rgb(0 0 0 / 20%), 0 1px 1px 0 rgb(0 0 0 / 14%), 0 1px 3px 0 rgb(0 0 0 / 12%)'
                      }}>
                        {decodeChannel(msg.type) == 'Email' ?
                          <div className={"text-overflow-ellipsis {min-height: 40px; & * {margin-bottom: 2px;}}"}
                               dangerouslySetInnerHTML={{__html: sanitizeHtml(JSON.parse(msg.content).html)}}></div> :
                          <div className="flex">{msg.content}</div>
                        }
                        <div className="flex align-content-end" style={{
                          width: '100%',
                          textAlign: 'right',
                          fontSize: '12px',
                          paddingTop: '15px',
                          color: '#C1C1C1'
                        }}>
                          <span className="flex-grow-1"></span>
                          <span
                            className="text-gray-600 mr-2">{decodeChannel(msg.type)}</span>
                          <Moment className="text-sm text-gray-600" date={t}
                                  format={'HH:mm'}></Moment>
                        </div>
                      </div>
                      <div className="flex flex-grow-1"></div>
                    </div>
                  </>
                }
                {
                  msg.direction == 1 &&
                  <>

                    {
                      (index === 0 || (index > 0 && messages[index - 1].direction !== messages[index].direction)) &&
                      <div className="w-full flex">
                        <div className="flex-grow-1"></div>
                        <div className="flex-grow-0 mb-1 pr-3">{msg.senderUsername.identifier}</div>
                      </div>
                    }

                    <div className="w-full flex">
                      <div className="flex-grow-1"></div>
                      <div className="flex-grow-0 flex-column p-3"
                           style={{background: '#C5EDCE', borderRadius: '5px'}}>
                        <div
                          className={"flex"}
                          dangerouslySetInnerHTML={{__html: sanitizeHtml(decodeChannel(msg.type) === "Email" ? JSON.parse(msg.content).html : msg.content)}}
                        ></div>
                        <div className="flex align-content-end" style={{
                          width: '100%',
                          textAlign: 'right',
                          fontSize: '12px',
                          paddingTop: '15px',
                          color: '#C1C1C1'
                        }}>
                          <span className="flex-grow-1"></span>
                          <span
                            className="text-gray-600 mr-2">{decodeChannel(msg.type)}</span>
                          <Moment className="text-sm text-gray-600" date={t}
                                  format={'HH:mm'}></Moment>
                        </div>
                      </div>
                    </div>
                  </>
                }
              </div>
            })
          }
        </div>
        <div id="chatWindowToScroll"></div>
        {/* TODO: This is the div that will be scrolled to smoothly on message send*/}
        <div id="chatWindowToScrollOnLoad"></div>
        {/* TODO: This is the div that will be scrolled to instantly on page load*/}
      </div>
      <div className="flex-grow-0 w-full p-3 pt-1">
        <div className="w-full h-full bg-white py-2" style={{
          border: 'solid 1px #E8E8E8',
          borderRadius: '8px',
          boxShadow: '0px 0px 40px rgba(0, 0, 0, 0.05)'
        }}>
          {currentChannel == "EMAIL" &&
            <div className="py-2">To: {showParticipants()} {addParticipantBox()}</div>}
          <div className="flex flex-grow-1">
            <InputTextarea className="w-full px-3 outline-none"
                           value={currentText}
                           onChange={(e) => setCurrentText(e.target.value)}
                           autoResize
                           rows={1}
                           placeholder={feedInitiator.firstName && `Message ${feedInitiator.firstName}...`}
                           onKeyPress={(e) => {
                             if (e.shiftKey && e.key === "Enter") {
                               return true
                             }
                             if (e.key === "Enter") {
                               handleSendMessage()
                               e.preventDefault();
                             }
                           }}
                           style={{
                             borderColor: "white", //Do not set as none!! It breaks InputTextarea autoResize
                             boxShadow: "none"
                           }}
            />
          </div>

          <div className="flex w-full mt-3">
            <div className="flex flex-grow-1">
              <div className="pl-1">
              </div>
              <Tooltip target=".disabled-button"/>
              <Tooltip target=".disabled-button2"/>
              <div className="disabled-button" data-pr-tooltip="Work in progress">
                <Button disabled={true} className='p-button-text'>
                  <FontAwesomeIcon icon={faSmile} style={{fontSize: '20px'}}/>
                </Button>
              </div>
              <div className="disabled-button2" data-pr-tooltip="Work in progress">
                <Button disabled={true} className='p-button-text'>
                  <FontAwesomeIcon icon={faPaperclip} style={{fontSize: '20px'}}/>
                </Button>
              </div>
            </div>
            {
              callingAllowed() && !props.inCall &&
              <div>
                <Button
                  onClick={() => props.handleCall(feedInitiator)}
                  tooltip={
                    `Call (${feedInitiator.phoneNumber})`
                  }
                  tooltipOptions={{position: 'top', showDelay: 200, hideDelay: 200}}
                  className='p-button-text mx-2 p-2'
                  style={{
                    border: 'solid 1px #E8E8E8',
                    borderRadius: '6px'
                  }}>
                  <FontAwesomeIcon icon={faPhone} style={{fontSize: '20px'}}/>
                </Button>
              </div>
            }

            <div className="flex flex-grow-0 mr-2">
              {/* TODO: Add Icon to left of reply button that changes based on chat type (dropdown??) */}
              <SplitButton
                model={sendButtonOptions}
                // disabled={sendButtonDisabled}
                onClick={() => handleSendMessage()}
                label={
                  `Reply (${currentChannel})`
                }
                className='p-button-text'
                style={{
                  background: 'var(--gray-color-1)',
                  border: 'solid 1px #E8E8E8',
                  borderRadius: '6px'
                }}
              >
                {/* <FontAwesomeIcon icon={faPaperPlane} className="mr-3" /> */}
              </SplitButton>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Chat;
