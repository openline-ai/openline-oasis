import {useCallback, useEffect, useRef, useState} from "react";
import * as React from "react";
import {Button} from "primereact/button";
import {faPaperclip, faPlay, faPhone, faPhoneSlash} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {InputText} from "primereact/inputtext";
import {useRouter} from "next/router";
import axios from "axios";
import {random} from "nanoid";
import {Dropdown} from "primereact/dropdown";
import Layout from "../../components/layout/layout";
import WebRTC from "./WebRTC";
import {IFrame} from "@stomp/stompjs";
import {configureStomp, useStomp} from "./useStomp";

export const Chat = ({user}: any) => {
    const router = useRouter();

    const {id} = router.query;
    const incomingMessage = useStomp();

    const messageWrapper:React.RefObject<HTMLDivElement> = useRef(null);


    const [currentUser, setCurrentUser] = useState({
        username: 'AgentSmith',
        firstName: 'Agent',
        lastName: 'Smith'
    });

    const [currentCustomer, setCurrentCustomer] = useState({
        username: 'customer1',
        firstName: 'John',
        lastName: 'Doe'
    });

    const [currentCompany, setCurrentCompany] = useState({
        name: 'Google'
    });

    function zeroPad(number: number) {
        if(number < 10) return '0' + number;
        return '' + number;
    }

    function monthConvert(number: number) {
        let months = ['Jan.', 'Feb.', 'Mar.', 'Apr.', 'May', 'June', 'July', 'Aug.', 'Sept.', 'Oct.', 'Nov.', 'Dec.'];
        return months[number-1];
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
            axios.get(`${process.env.NEXT_PUBLIC_BE_PATH}/case/${id}/item`)
                .then(res => {
                    setMessageList(res.data);
                });
            axios.get(`${process.env.NEXT_PUBLIC_BE_PATH}/case/${id}`)
                .then(res => {
                    setCurrentCustomer({username: res.data.userName, firstName: "John", lastName: "doe"});

                });
        }
    }, [id]);

    useEffect(() => {
        configureStomp(`${process.env.NEXT_PUBLIC_STOMP_WEBSOCKET_PATH}/websocket`, `/queue/new-case-item/${id}`);
    }, [id])

    const refreshCredentials = () => {
        axios.get(`${process.env.NEXT_PUBLIC_BE_PATH}/call_credentials/?service=sip&username=`+currentUser.username + "@agent.openline.ai")
            .then(res => {
                console.error("Got a key: " + JSON.stringify(res.data));
                if(webrtc.current?._ua) {
                    webrtc.current?.stopUA();
                }
                webrtc.current?.setCredentials(res.data.username, res.data.password,
                    () =>{webrtc.current?.startUA()});
                setTimeout(() => {refreshCredentials()}, (res.data.ttl*3000)/4);
            });
    }
    useEffect(() => {
        refreshCredentials();
    }, []);


    useEffect(() => {
        setMessages(messageList.map((msg: any) => {
            console.log("Have a message:\n" + JSON.stringify(msg));
            let lines = msg.message.split('\n');

            let filtered:string[] = lines.filter(function (line:string) {
                return line.indexOf('>') != 0;
            });
            msg.message = filtered.join('\n').trim();
            let year = msg.createdDate[0];
            let month = monthConvert(msg.createdDate[1]);
            let day = msg.createdDate[2];
            let hour = zeroPad(msg.createdDate[3]);
            let minute = zeroPad(msg.createdDate[4]);

            return (<div key={msg.id} style={{
                display: 'block',
                width: 'auto',
                maxWidth: '100%',
                wordBreak: 'break-all',
                padding: '10px',
                margin: '0px 5px'
            }}>
                {msg.direction === 'INBOUND' &&
                    <div style={{textAlign: 'left'}}>
                        <div style={{
                            fontSize: '10px',
                            marginBottom: '10px'
                        }}>{currentCustomer.username}&nbsp;-&nbsp;{msg.channel}&nbsp;-&nbsp;{day},&nbsp;{month}&nbsp;{year}&nbsp;{hour}:{minute}</div>
                        <span style={{whiteSpace: 'pre-wrap', background: '#bbbbbb', lineHeight: '27px', borderRadius: '3px', padding: '7px 10px'}}>
                    <span style={{}}>{msg.message}</span><span style={{marginLeft: '10px'}}></span>
                    </span>
                    </div>
                }
                {msg.direction === 'OUTBOUND' &&
                    <div style={{textAlign: 'right'}}>
                        <div style={{
                            fontSize: '10px',
                            lineHeight: '16px',
                            marginBottom: '10px'
                        }}>{currentUser.firstName}&nbsp;{currentUser.lastName}&nbsp;-&nbsp;{day},&nbsp;{month}&nbsp;{year}&nbsp;{hour}:{minute}</div>
                        <span style={{whiteSpace: 'pre-wrap', background: '#bbbbbb', lineHeight: '27px', borderRadius: '3px', padding: '7px 10px'}}>
                            <span style={{}}>{msg.message}</span><span style={{marginLeft: '10px'}}>{hour}:{minute}</span>
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
        console.log("Got websocket!" + JSON.stringify(incomingMessage));
        if (incomingMessage && Object.keys(incomingMessage).length !== 0 && incomingMessage.body.length > 0) {
            handleWebsocketMessage(JSON.parse(incomingMessage.body));
        }
    }, [incomingMessage]);

    const webrtc:React.RefObject<WebRTC> = useRef<WebRTC>(null);

    const handleCall = () => {
        //setInCall(true);
            let user = currentCustomer.username;
            const regex = /.*<(.*)>/;
            const matches = user.match(regex);
            if(matches) {
                user = matches[1];
            }
            webrtc.current?.makeCall("sip:" + user);
    }
    const hangupCall = () => {
        setInCall(false);
        webrtc.current?.hangupCall();

    }

    const handleSendMessage = () => {
        axios.post(`${process.env.NEXT_PUBLIC_BE_PATH}/case/${id}/item`, {
            source: 'WEB',
            direction: 'OUTBOUND',
            channel: currentChannel,
            userName: currentCustomer.username,
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
            id: msg.id,
            direction: msg.direction,
            message: msg.message,
            createdDate: msg.createdDate,
            channel: msg.channel
        };
        console.log("Adding message: " + JSON.stringify(newMsg));
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
                        updateCallState={(state: boolean)=>setInCall(state)}
                    autoStart={false}

                /> }
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

                    <Button disabled={sendButtonDisabled} onClick={() => handleSendMessage()} className='p-button-text'>
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
