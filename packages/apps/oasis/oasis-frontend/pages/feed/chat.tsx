import * as React from "react";
import {useEffect, useState} from "react";
import {Button} from "primereact/button";
import {SplitButton} from 'primereact/splitbutton';
import {faPaperclip, faPhone, faSmile} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {InputTextarea} from "primereact/inputtextarea";
import axios from "axios";
import useWebSocket from "react-use-websocket";
import {loggedInOrRedirectToLogin} from "../../utils/logged-in";
import {getSession, useSession} from "next-auth/react";
import {gql, GraphQLClient} from "graphql-request";
import {ProgressSpinner} from "primereact/progressspinner";
import {Tooltip} from 'primereact/tooltip';
import Moment from "react-moment";
import {FeedItem} from "../../model/feed-item";
import {toast} from "react-toastify";
import {ConversationItem} from "../../model/conversation-item";

interface ChatProps {
    feedId: string;
    inCall: boolean;

    handleCall(feedInitiator: any): void;
}

export const Chat = (props: ChatProps) => {
    const client = new GraphQLClient(`/customer-os-api/query`);

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
    const [sendButtonDisabled, setSendButtonDisabled] = useState(false);
    const [messages, setMessages] = useState([] as ConversationItem[]);
    const {data: session, status} = useSession();

    const [loadingMessages, setLoadingMessages] = useState(false)

    useEffect(() => {
        if (props.feedId) {
            setLoadingMessages(true);
            setCurrentText('');

            axios.get(`/oasis-api/feed/${props.feedId}`)
            .then(res => {
                const feedItem = res.data as FeedItem;

                if (feedItem.initiatorType === 'CONTACT') {

                    const query = gql`query GetContactDetails($email: String!) {
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

                    client.request(query, {email: feedItem.initiatorUsername}).then((response: any) => {
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

                    //TODO move initiator in index
                } else if (feedItem.initiatorUsername === 'USER') {

                    const query = gql`query GetUserById {
                        user(id: "${feedItem.initiatorUsername}") {
                            id
                            firstName
                            lastName
                        }
                    }`

                    client.request(query).then((response: any) => {
                        if (response.user) {
                            setFeedInitiator({
                                loaded: true,
                                firstName: response.user.firstName,
                                lastName: response.user.lastName,
                                email: response.user.emails[0]?.email ?? undefined,
                                phoneNumber: response.user.phoneNumbers[0]?.e164 ?? undefined //TODO user doesn't have phone in backend
                            });
                        } else {
                            //TODO log on backend
                            toast.error("There was a problem on our side and we are doing our best to solve it!");
                        }
                    }).catch(reason => {
                        //TODO log on backend
                        toast.error("There was a problem on our side and we are doing our best to solve it!");
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
        axios.post(`/oasis-api/feed/${props.feedId}/item`, {
            channel: currentChannel,
            username: session?.user?.email,
            message: currentText
        }).then(res => {
            console.log(res)
            if (res.data) {
                setMessages((messageList: any) => [...messageList, res.data]);
                setCurrentText('');
            }
        }).catch(reason => {
            //todo log on backend
            toast.error("There was a problem on our side and we are doing our best to solve it!");
        });
    };

    const handleWebsocketMessage = function (msg: any) {
        let newMsg = {
            content: msg.message,
            username: msg.username,
            channel: 1,
            time: msg.time,
            id: msg.id,
            direction: msg.direction == "OUTBOUND" ? 1 : 0,
            contact: {},
        };

        setMessages((messageList: any) => [...messageList, newMsg]);
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
                                messages.map((msg: ConversationItem, index: any) => {
                                    let lines = msg.content.split('\n');

                                    let filtered: string[] = lines.filter(function (line: string) {
                                        return line.indexOf('>') != 0;
                                    });
                                    msg.content = filtered.join('\n').trim();

                                    var t = new Date(1970, 0, 1);
                                    t.setSeconds(msg.time.seconds);

                                    return <div key={msg.id} className='flex flex-column mb-3'>
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
                                                            <div className="flex">{msg.content}</div>
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
                                                                <div className="flex-grow-0 mb-1 pr-3">To be added</div>
                                                            </div>
                                                    }

                                                    <div className="w-full flex">
                                                        <div className="flex-grow-1"></div>
                                                        <div className="flex-grow-0 flex-column p-3"
                                                             style={{background: '#C5EDCE', borderRadius: '5px'}}>
                                                            <div className="flex">{msg.content}</div>
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

                        <div className="flex flex-grow-1">
                            <InputTextarea className="w-full px-3 outline-none"
                                           value={currentText}
                                           onChange={(e) => setCurrentText(e.target.value)}
                                           autoResize
                                           rows={1}
                                           placeholder={
                                                   feedInitiator.firstName &&
                                                   `Message ${feedInitiator.firstName}...`
                                           }
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

export async function getServerSideProps(context: any) {
    return loggedInOrRedirectToLogin(await getSession(context));
}

export default Chat;
