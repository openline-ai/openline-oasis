import {styles} from "./styles";
import {Button, Input, MessageList} from "react-chat-elements";
import React, {useEffect, useRef, useState} from "react";
import useWebSocket from 'react-use-websocket';

let clearRef = () => {
}

function useForceUpdate() {
    const [value, setValue] = useState(0)
    return () => setValue(() => value + 1)
}

interface ChatEngineProps {
    user: string,
    apikey: string,
    httpServerPath: string
    wsServerPath: string
}

export default function ChatEngine(props: ChatEngineProps) {
    const [datasource, setDatasource] = useState([] as any);
    const [messageText, setMessageText] = useState<string>("");
    const messageListReferance = useRef()
    const inputReferance = useRef();
    const forceUpdate = useForceUpdate();

    const {lastMessage} = useWebSocket(`${props.wsServerPath}/ws/${props.user}`, {
        onOpen: () => console.log('Websocket opened'),
        //Will attempt to reconnect on all close events, such as server shutting down
        shouldReconnect: (closeEvent) => true,
    });

    useEffect(() => {
        if (lastMessage && Object.keys(lastMessage).length !== 0 && lastMessage.data.length > 0) {
            let parsed = JSON.parse(lastMessage.data);
            setDatasource((prevDatasource: any) => [...prevDatasource, {
                position: 'left',
                type: 'text',
                text: parsed.message,
                date: new Date(),
            }]);
        }
    }, [lastMessage])

    const handleMessage = (msg: string) => {
        fetch(`${props.httpServerPath}/webchat/`, {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'follow',
            referrerPolicy: 'no-referrer',
            headers: {
                "Content-Type": "application/json",
                'WebChatApiKey': `${props.apikey}`
            },
            body: JSON.stringify({"username": `${props.user}`, "message": msg}),
        }).then((response) => {
            if (!response.ok) {
                throw Error(response.statusText);
            }
            return response.json();
        }).then(response => {
            let stringify = JSON.stringify(response);
            setDatasource((prevDatasource: any) => [...prevDatasource, {
                position: 'right',
                type: 'text',
                text: msg,
                date: new Date(),
            }]);
            clearRef();
            forceUpdate();
            console.log(stringify)
        }).catch(error => {
            console.log(error);
        });
    }

    return (
        <div style={{...styles.divStyle}}>
            <MessageList
                referance={messageListReferance}
                className='message-list'
                toBottomHeight={'100%'}
                dataSource={datasource}
                lockable={true}
                downButton={false}
                downButtonBadge={10}
                sendMessagePreview={true}
            />
            <Input
                referance={inputReferance}
                clear={(clear: () => {}) => (clearRef = clear)}
                placeholder='Write your message here.'
                multiline={true}
                maxHeight={50}
                onChange={(e: any) => setMessageText(e.target.value)}
                onKeyPress={(e) => {
                    if (e.shiftKey && e.key === "Enter") {
                        return true
                    }
                    if (e.key === "Enter") {
                        clearRef();
                        handleMessage(messageText);
                    }
                }}
                rightButtons={<Button text={"Send"} onClick={() => handleMessage(messageText)}
                                      title="Send"/>}
            />
        </div>
    )
}