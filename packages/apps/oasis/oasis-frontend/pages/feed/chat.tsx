import * as React from "react";
import {useCallback, useEffect, useRef, useState} from "react";
import {Button} from "primereact/button";
import {faPhone, faPhoneSlash, faPlay} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {InputText} from "primereact/inputtext";
import {useRouter} from "next/router";
import axios from "axios";
import {Dropdown} from "primereact/dropdown";
import Layout from "../../components/layout/layout";
import WebRTC from "./WebRTC";
import useWebSocket from "react-use-websocket";

export const Chat = ({user}: any) => {
    const router = useRouter();

    const {id} = router.query;

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

    const [currentCustomer, setCurrentCustomer] = useState({
        contactId: 'customer1',
        firstName: 'John',
        lastName: 'Doe',
        lastMailAddress: ''
    });

    const [currentCompany, setCurrentCompany] = useState({
        name: 'Google'
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

    const [currentChannel, setCurrentChannel] = useState('CHAT');
    const [currentText, setCurrentText] = useState('');
    const [attachmentButtonHidden, setAttachmentButtonHidden] = useState(false);
    const [sendButtonDisabled, setSendButtonDisabled] = useState(false);
    const [inCall, setInCall] = useState(false);
    const [messageList, setMessageList] = useState([] as any);
    const [messages, setMessages] = useState([] as any);

    useEffect(() => {
        if (id) {

            axios.get(`/server/feed/${id}`)
                .then(res => {
                    setCurrentCustomer({
                        contactId: res.data.contactId,
                        firstName: res.data.firstName,
                        lastName: res.data.lastName,
                        lastMailAddress: ''
                    });
                    axios.get(`/server/feed/${id}/item`)
                    .then(res => {
                        setMessageList(res.data);
                    });

                });
        }
    }, [id]);

    const refreshCredentials = () => {
        axios.get(`/server/call_credentials/?service=sip&username=` + currentUser.username + "@agent.openline.ai")
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
    useEffect(() => {
        refreshCredentials();
    }, []);


    useEffect(() => {
        setMessages(messageList.map((msg: any) => {
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

            if (msg.channel == 0 || msg.channel == 1) {
                setCurrentCustomer({
                    contactId: currentCustomer.contactId,
                    firstName: currentCustomer.firstName,
                    lastName: currentCustomer.lastName,
                    lastMailAddress: msg.username
                });
            }

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
                        }}>{msg.username}&nbsp;-&nbsp;{decodeChannel(msg.channel)}&nbsp;-&nbsp;{day},&nbsp;{month}&nbsp;{year}&nbsp;{hour}:{minute}</div>
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
    }, [messageList]);

    //when a new message appears, scroll to the end of container
    useEffect(() => {
        // @ts-ignore
        messageWrapper?.current?.scrollIntoView({behavior: "smooth"});
    }, [messages]);

    //when the user types, we hide the buttons
    useEffect(() => {
        setAttachmentButtonHidden(currentText !== '')
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
        let user = currentCustomer.lastMailAddress;
        const regex = /.*<(.*)>/;
        const matches = user.match(regex);
        if (matches) {
            user = matches[1];
        }
        webrtc.current?.makeCall("sip:" + user);
    }
    const hangupCall = () => {
        setInCall(false);
        webrtc.current?.hangupCall();

    }

    const handleSendMessage = () => {
        axios.post(`/server/feed/${id}/item`, {
            source: 'WEB',
            direction: 'OUTBOUND',
            channel: currentChannel,
            username: currentCustomer.lastMailAddress,
            message: currentText
        })
            .then(res => {
                setMessageList((messageList: any) => [...messageList, res.data]);
                setCurrentText('');
            });
    };


    const getTypingIndicator = useCallback(
        () => {
            return undefined;
        }, [],
    );

    const handleWebsocketMessage = function (msg: any) {
        let newMsg = {
            message: msg.message,
            username: msg.username,
            channel: 1,
            time: msg.time,
            id: msg.id,
            contact: {},
        };

        setMessageList((messageList: any) => [...messageList, newMsg]);
    }

    return (
        <>
            <Layout>
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
                            from={"sip:" + currentUser.username + "@agent.openline.ai"}
                            updateCallState={(state: boolean) => setInCall(state)}
                            autoStart={false}

                        />}
                    {messages}
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
                                         {process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL &&
                                             <FontAwesomeIcon icon={faPhone} style={{color: 'black'}}/>
                                         }
                    </Button>
                    </span>
                    <span hidden={!inCall}>
                            <Button onClick={() => hangupCall()} className='p-button-text' hidden={!inCall}>
                            {process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL &&
                                <FontAwesomeIcon icon={faPhoneSlash} style={{color: 'black'}}/>
                            }
                            </Button>
                    </span>
                </div>

            </Layout>
        </>
    );

}

export default Chat;
