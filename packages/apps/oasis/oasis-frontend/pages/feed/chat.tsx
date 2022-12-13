import * as React from "react";
import {useEffect, useRef, useState} from "react";
import {Button} from "primereact/button";
import {faPhone, faPhoneSlash, faPlay, faRightLeft} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {InputText} from "primereact/inputtext";
import axios from "axios";
import {Dropdown} from "primereact/dropdown";
import WebRTC from "./WebRTC";
import useWebSocket from "react-use-websocket";
import {loggedInOrRedirectToLogin} from "../../utils/logged-in";
import {getSession, useSession} from "next-auth/react";
import {gql, GraphQLClient} from "graphql-request";


export const Chat = ({id}: { id: string | string[] | undefined }) => {
    const client = new GraphQLClient(`${process.env.NEXT_PUBLIC_CUSTOMER_OS_API_PATH}/query`);

    const {lastMessage} = useWebSocket(`${process.env.NEXT_PUBLIC_WEBSOCKET_PATH}/${id}`, {
        onOpen: () => console.log('Websocket opened'),
        //Will attempt to reconnect on all close events, such as server shutting down
        shouldReconnect: (closeEvent) => true,
    })

    const messageWrapper: React.RefObject<HTMLDivElement> = useRef(null);

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

    function zeroPad(number: number) {
        if (number < 10) return '0' + number;
        return '' + number;
    }

    function monthConvert(number: number) {
        let months = ['Jan.', 'Feb.', 'Mar.', 'Apr.', 'May', 'June', 'July', 'Aug.', 'Sept.', 'Oct.', 'Nov.', 'Dec.'];
        return months[number - 1];
    }

    function decodeChannel(channel: number) {
        switch (channel) {
            case 0:
                return "CHAT";
            case 1:
                return "MAIL";
            case 2:
                return "WHATSAPP";
            case 3:
                return "FACEBOOK";
            case 4:
                return "TWITTER";
            case 5:
                return "VOICE";
        }
        return "CHAT";
    }

    function callingAllowed() {
        return process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL &&
                (contact.phoneNumber || contact.email == "echo@oasis.openline.ai");
    }

    const [currentChannel, setCurrentChannel] = useState('CHAT');
    const [currentText, setCurrentText] = useState('');
    const [sendButtonDisabled, setSendButtonDisabled] = useState(false);
    const [inCall, setInCall] = useState(false);
    const [messageList, setMessageList] = useState([] as any);
    const [messages, setMessages] = useState([] as any);
    const {data: session, status} = useSession();

    const [loadingMessages, setLoadingMessages] = useState(false)

    useEffect(() => {
        if (id) {
            setLoadingMessages(true);
            console.log('load feed data');
            axios.get(`/server/feed/${id}`)
            .then(res => {
                console.log('load feed data completed');
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
                        console.log('current contact loaded');
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

            axios.get(`/server/feed/${id}/item`)
                    .then(res => {
                        setMessageList(res.data ?? []);
                    }).catch((reason: any) => {
                //TODO error
            });
        }
    }, [id]);

    useEffect(() => {

        const refreshCredentials = () => {
            axios.get(`/server/call_credentials?service=sip&username=` + session?.user?.email)
                    .then(res => {
                        console.error("Got a key: " + JSON.stringify(res.data));
                        if (webrtc.current?._ua) {
                            webrtc.current?.stopUA();
                        }
                        webrtc.current?.setCredentials(res.data.username, res.data.password,
                                () => {
                                    webrtc.current?.startUA()
                                });
                        setTimeout(() => {
                            refreshCredentials()
                        }, (res.data.ttl * 3000) / 4);
                    });
        }
        if (session?.user?.email) {
            refreshCredentials();
        }
    }, [session]);

    useEffect(() => {
        setMessages(messageList?.map((msg: any) => {
            console.log("Have a message:\n" + JSON.stringify(msg));
            let lines = msg.message.split('\n');

            let filtered: string[] = lines.filter(function (line: string) {
                return line.indexOf('>') != 0;
            });
            msg.message = filtered.join('\n').trim();
            let t = new Date(Date.UTC(1970, 0, 1));
            t.setUTCSeconds(msg.time.seconds);
            let year = t.getFullYear();
            let month = monthConvert(t.getMonth() + 1);
            let day = t.getDate();
            let hour = zeroPad(t.getHours());
            let minute = zeroPad(t.getMinutes());

            return (<div key={msg.id} style={{
                display: 'block',
                width: 'auto',
                maxWidth: '100%',
                wordBreak: 'break-all',
                padding: '10px',
                margin: '0px 5px'
            }}>
                {!msg.direction &&
                        <div style={{textAlign: 'left'}}>
                            <div style={{
                                fontSize: '10px',
                                marginBottom: '10px'
                            }}>

                                {/*{contact.firstName && contact.firstName + ' ' + contact.lastName}*/}
                                {/*{!contact.firstName && contact.email && contact.email}*/}
                                {/*{!contact.firstName && !contact.email && contact.phoneNumber}*/}

                                {decodeChannel(msg.channel)}&nbsp;-&nbsp;{day},&nbsp;{month}&nbsp;{year}&nbsp;{hour}:{minute}</div>
                            <span style={{
                                whiteSpace: 'pre-wrap',
                                background: '#bbbbbb',
                                lineHeight: '27px',
                                borderRadius: '3px',
                                padding: '7px 10px'
                            }}>
                    <span style={{}}>{msg.message}</span><span style={{marginLeft: '10px'}}></span>
                    </span>
                        </div>
                }
                {msg.direction == 1 &&
                        <div style={{textAlign: 'right'}}>
                            <div style={{
                                fontSize: '10px',
                                lineHeight: '16px',
                                marginBottom: '10px'
                            }}>{currentUser.firstName}&nbsp;{currentUser.lastName}&nbsp;-&nbsp;{day},&nbsp;{month}&nbsp;{year}&nbsp;{hour}:{minute}</div>
                            <span style={{
                                whiteSpace: 'pre-wrap',
                                background: '#bbbbbb',
                                lineHeight: '27px',
                                borderRadius: '3px',
                                padding: '7px 10px'
                            }}>
                            <span style={{}}>{msg.message}</span><span style={{marginLeft: '10px'}}></span>
                        </span>
                        </div>
                }

            </div>);
        }));
        console.log('messages list loaded')
    }, [messageList]);

    //when a new message appears, scroll to the end of container
    useEffect(() => {
        // @ts-ignore
        messageWrapper?.current?.scrollIntoView({behavior: "smooth"});
        setLoadingMessages(false);
        console.log('messages loaded')
    }, [messages]);

    //when the user types, we hide the buttons
    useEffect(() => {
        setSendButtonDisabled(currentText === '')
    }, [currentText]);

    useEffect(() => {
        if (lastMessage && Object.keys(lastMessage).length !== 0 && lastMessage.data.length > 0) {
            handleWebsocketMessage(JSON.parse(lastMessage.data));
        }
    }, [lastMessage]);

    const webrtc: React.RefObject<WebRTC> = useRef<WebRTC>(null);

    const handleCall = () => {
        //setInCall(true);
        let user = '';
        if (contact.phoneNumber) {
            user = contact.phoneNumber + "@oasis.openline.ai";
        } else {
            user = contact.email;
            const regex = /.*<(.*)>/;
            const matches = user.match(regex);
            if (matches) {
                user = matches[1];
            }
        }
        webrtc.current?.makeCall("sip:" + user);
    }
    const hangupCall = () => {
        setInCall(false);
        webrtc.current?.hangupCall();

    }

    const showTransfer = () => {
        webrtc.current?.showTransfer();

    }

    const handleSendMessage = () => {
        axios.post(`/server/feed/${id}/item`, {
            source: 'WEB',
            direction: 'OUTBOUND',
            channel: currentChannel,
            username: contact.email,
            message: currentText
        })
                .then(res => {
                    setMessageList((messageList: any) => [...messageList, res.data]);
                    setCurrentText('');
                });
    };

    const handleWebsocketMessage = function (msg: any) {
        let newMsg = {
            message: msg.message,
            username: msg.username,
            channel: 1,
            time: msg.time,
            id: msg.id,
	    direction: msg.direction == "OUTBOUND"?1:0,
            contact: {},
        };

        setMessageList((messageList: any) => [...messageList, newMsg]);
    }

    return (
            <div className='w-full h-full'>
                <div style={{
                    width: '100%',
                    height: 'calc(100% - 100px)',
                    overflowX: 'hidden',
                    overflowY: 'auto'
                }}>
                    {process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL &&
                            <WebRTC
                                    ref={webrtc}
                                    websocket={`${process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL}`}
                                    from={"sip:" + session?.user?.email}
                                    updateCallState={(state: boolean) => setInCall(state)}
                                    autoStart={false}

                            />}
                    {
                            loadingMessages &&
                            <div>Loading</div>
                    }

                    {
                            !loadingMessages &&
                            messages
                    }

                    <div ref={messageWrapper}></div>
                </div>
                <div style={{width: '100%', height: '100px'}}>

                    <InputText style={{width: 'calc(100% - 150px)'}} value={currentText}
                               onChange={(e) => setCurrentText(e.target.value)}/>

                    <Dropdown optionLabel="label" value={currentChannel} options={[
                        {
                            label: 'Web chat',
                            value: 'CHAT'
                        },
                        {
                            label: 'Email',
                            value: 'EMAIL'
                        },

                    ]} onChange={(e) => setCurrentChannel(e.value)}/>

                    <Button disabled={sendButtonDisabled} onClick={() => handleSendMessage()}
                            className='p-button-text'>
                        <FontAwesomeIcon icon={faPlay} style={{color: 'black'}}/>
                    </Button>

                    {/*{*/}
                    {/*    !attachmentButtonHidden &&*/}
                    {/*    <>*/}
                    {/*        <Button onClick={() => fileUploadInput?.current.click()} className='p-button-text'>*/}
                    {/*            <FontAwesomeIcon icon={faPaperclip} style={{color: 'black'}}/>*/}
                    {/*        </Button>*/}
                    {/*        <input ref={fileUploadInput} type="file" name="file" style={{display: 'none'}}/>*/}
                    {/*    </>*/}
                    {/*}*/}

                    <span hidden={inCall}>
                    <Button onClick={() => handleCall()} className='p-button-text' hidden={inCall}>
                                         {callingAllowed() &&
                                                 <FontAwesomeIcon icon={faPhone} style={{color: 'black'}}/>
                                         }
                    </Button>
                    </span>
                    <span hidden={!inCall}>
                            <Button onClick={() => hangupCall()} className='p-button-text' hidden={!inCall}>
                            {callingAllowed() &&
                                    <FontAwesomeIcon icon={faPhoneSlash} style={{color: 'black'}}/>
                            }
                            </Button>
                            <Button onClick={() => showTransfer()} className='p-button-text' hidden={!inCall}>
                            {callingAllowed() &&
                                    <FontAwesomeIcon icon={faRightLeft} style={{color: 'black'}}/>
                            }
                            </Button>
                    </span>
                </div>
            </div>
    );

}

export async function getServerSideProps(context: any) {
    return loggedInOrRedirectToLogin(await getSession(context));
}

export default Chat;
