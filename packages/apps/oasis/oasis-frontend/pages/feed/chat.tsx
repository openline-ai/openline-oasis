import * as React from "react";
import {useCallback, useEffect, useRef, useState} from "react";
import {Button} from "primereact/button";
import {
    faPaperclip,
    faPaperPlane,
    faPhone,
    faPhoneSlash,
    faRightLeft,
    faSmile
} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {InputText} from "primereact/inputtext";
import axios from "axios";
import {Dropdown} from "primereact/dropdown";
import useWebSocket from "react-use-websocket";
import {loggedInOrRedirectToLogin} from "../../utils/logged-in";
import {getSession, useSession} from "next-auth/react";
import {gql, GraphQLClient} from "graphql-request";
import {ProgressSpinner} from "primereact/progressspinner";
import {Tooltip} from "primereact/tooltip";
import Moment from "react-moment";

interface ChatProps {
    feedId: string;
    inCall: boolean;

    handleCall(contact: any): void;

    hangupCall(): void;

    showTransfer(): void;
}

export const Chat = (props: ChatProps) => {
    const client = new GraphQLClient(`${process.env.NEXT_PUBLIC_CUSTOMER_OS_API_PATH}/query`);

    const {lastMessage} = useWebSocket(`${process.env.NEXT_PUBLIC_WEBSOCKET_PATH}/${props.feedId}`, {
        onOpen: () => console.log('Websocket opened'),
        //Will attempt to reconnect on all close events, such as server shutting down
        shouldReconnect: (closeEvent) => true,
    })

    const [currentUser, setCurrentUser] = useState({
        username: 'AgentSmith',
        firstName: 'Agent',
        lastName: 'Smith'
    });

    const [contact, setContact] = useState({
        firstName: '',
        lastName: '',
        email: '',
        phoneNumber: '',
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
        return process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL && (contact.phoneNumber || contact.email == "echo@oasis.openline.ai");
    }

    const [currentChannel, setCurrentChannel] = useState('CHAT');
    const [currentText, setCurrentText] = useState('');
    const [sendButtonDisabled, setSendButtonDisabled] = useState(false);
    const [messages, setMessages] = useState([] as any);
    const {data: session, status} = useSession();

    const [loadingMessages, setLoadingMessages] = useState(false)

    useEffect(() => {
        if (props.feedId) {
            setLoadingMessages(true);
            setCurrentText('');

            axios.get(`/server/feed/${props.feedId}`)
            .then(res => {
                const query = gql`query GetContactDetails($id: ID!) {
                    contact(id: $id) {
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

                client.request(query, {id: res.data.contactId}).then((response: any) => {
                    if (response.contact) {
                        setContact({
                            firstName: response.contact.firstName,
                            lastName: response.contact.lastName,
                            email: response.contact.emails[0]?.email ?? undefined,
                            phoneNumber: response.contact.phoneNumbers[0]?.e164 ?? undefined
                        });
                    } else {
                        //TODO error
                    }
                }).catch(reason => {
                    //TODO error
                });

            }).catch((reason: any) => {
                //TODO error
            });

            axios.get(`/server/feed/${props.feedId}/item`)
                    .then(res => {
                        setMessages(res.data ?? []);
                    }).catch((reason: any) => {
                //TODO error
            });
        }
    }, [props.feedId]);

    //when a new message appears, scroll to the end of container
    useEffect(() => {
        setLoadingMessages(false);
    }, [messages]);

    useEffect(() => {
        if (!loadingMessages) {
            const element = document.getElementById('chatWindowToScroll')
            if (element) {
                element.scrollIntoView({behavior: 'smooth'})
            }
        }
    }, [loadingMessages, messages])

    useEffect(() => {
        setSendButtonDisabled(currentText === '')
    }, [currentText]);

    useEffect(() => {
        if (lastMessage && Object.keys(lastMessage).length !== 0 && lastMessage.data.length > 0) {
            handleWebsocketMessage(JSON.parse(lastMessage.data));
        }
    }, [lastMessage]);

    const handleSendMessage = () => {
        axios.post(`/server/feed/${props.feedId}/item`, {
            source: 'WEB',
            direction: 'OUTBOUND',
            channel: currentChannel,
            username: contact.email,
            message: currentText
        }).then(res => {
            setMessages((messageList: any) => [...messageList, res.data]);
            setCurrentText('');
        }).catch(reason => {
            //TODO error
        });
    };

    const handleWebsocketMessage = function (msg: any) {
        let newMsg = {
            message: msg.message,
            username: msg.username,
            channel: 1,
            time: msg.time,
            id: msg.id,
            direction: msg.direction == "OUTBOUND" ? 1 : 0,
            contact: {},
        };

        setMessages((messageList: any) => [...messageList, newMsg]);
    }

    return (
            <div className='flex flex-column w-full h-full'>
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
                                messages.map((msg: any, index: any) => {
                                    let lines = msg.message.split('\n');

                                    let filtered: string[] = lines.filter(function (line: string) {
                                        return line.indexOf('>') != 0;
                                    });
                                    msg.message = filtered.join('\n').trim();

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
                                                                        contact.firstName && contact.lastName &&
                                                                        <>{contact.firstName} {contact.lastName}</>
                                                                }
                                                                {
                                                                        !contact.firstName && !contact.lastName &&
                                                                        <>{contact.email}</>
                                                                }
                                                            </div>
                                                    }

                                                    <div className="flex">
                                                        <div className="flex flex-column flex-grow-0 p-3" style={{
                                                            background: 'white',
                                                            borderRadius: '5px',
                                                            boxShadow: '0 2px 1px -1px rgb(0 0 0 / 20%), 0 1px 1px 0 rgb(0 0 0 / 14%), 0 1px 3px 0 rgb(0 0 0 / 12%)'
                                                        }}>
                                                            <div className="flex">{msg.message}</div>
                                                            <div className="flex align-content-end" style={{
                                                                width: '100%',
                                                                textAlign: 'right',
                                                                fontSize: '12px',
                                                                paddingTop: '15px',
                                                                color: '#C1C1C1'
                                                            }}>
                                                                <span className="flex-grow-1"></span>
                                                                <span className="text-gray-600 mr-2">{decodeChannel(msg.channel)}</span>
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
                                                                <div className="flex-grow-0 mb-1 pr-3">Dummy user</div>
                                                            </div>
                                                    }

                                                    <div className="w-full flex">
                                                        <div className="flex-grow-1"></div>
                                                        <div className="flex-grow-0 flex-column p-3"
                                                             style={{background: '#C5EDCE', borderRadius: '5px'}}>
                                                            <div className="flex">{msg.message}</div>
                                                            <div className="flex align-content-end" style={{
                                                                width: '100%',
                                                                textAlign: 'right',
                                                                fontSize: '12px',
                                                                paddingTop: '15px',
                                                                color: '#C1C1C1'
                                                            }}>
                                                                <span className="flex-grow-1"></span>
                                                                <span className="text-gray-600 mr-2">{decodeChannel(msg.channel)}</span>
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
                </div>
                <div className="flex-grow-0 w-full p-5">

                    <div className="w-full h-full bg-white p-5" style={{
                        border: 'solid 1px #E8E8E8',
                        borderRadius: '7px',
                        boxShadow: '0px 0px 40px rgba(0, 0, 0, 0.05)'
                    }}>

                        <Dropdown
                                className="border-none mb-3"
                                style={{width: '120px'}}
                                optionLabel="label"
                                value={currentChannel}
                                onChange={(e) => setCurrentChannel(e.value)}
                                options={[
                                    {
                                        label: 'Web chat',
                                        value: 'CHAT'
                                    },
                                    {
                                        label: 'Email',
                                        value: 'EMAIL'
                                    },
                                ]}/>

                        <div className="flex flex-grow-1">
                            <InputText className="w-full" value={currentText}
                                       onChange={(e) => setCurrentText(e.target.value)}
                                       onKeyPress={(e) => {
                                           if (e.shiftKey && e.key === "Enter") {
                                               return true
                                           }
                                           if (e.key === "Enter") {
                                               handleSendMessage()
                                           }
                                       }}/>
                        </div>

                        <div className="flex w-full mt-3">

                            <div className="flex flex-grow-1">

                                {
                                        callingAllowed() && !props.inCall &&
                                        <div>
                                            <Button onClick={() => props.handleCall(contact)} className='p-button-text'>
                                                <FontAwesomeIcon icon={faPhone} style={{fontSize: '20px'}}/>
                                            </Button>
                                        </div>
                                }

                                {
                                        callingAllowed() && props.inCall &&
                                        <div>
                                            <Button onClick={() => props.hangupCall()} className='p-button-text'>
                                                <FontAwesomeIcon icon={faPhoneSlash} style={{fontSize: '20px'}}/>
                                            </Button>
                                            <Button onClick={() => props.showTransfer()} className='p-button-text'>
                                                <FontAwesomeIcon icon={faRightLeft} style={{fontSize: '20px'}}/>
                                            </Button>
                                        </div>
                                }

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

                            <div className="flex flex-grow-0">
                                <Button disabled={sendButtonDisabled} onClick={() => handleSendMessage()}
                                        className='p-button-text'>
                                    <FontAwesomeIcon icon={faPaperPlane} className="mr-3"/>Reply
                                </Button>
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
