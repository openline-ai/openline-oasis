import React, {useState} from "react";
import EmailForm from "./EmailForm";
import {styles} from "./styles";
import ChatEngine from "./ChatEngine";

interface WebChatWindowProps {
    visible: boolean
    apikey: string,
    httpServerPath: string
    wsServerPath: string
}


export default function SupportWindow(props: WebChatWindowProps) {
    const [user, setUser] = useState<string>("")

    return (
        <div
            className='transition-5'
            style={{
                ...styles.supportWindow,
                ...{opacity: props.visible ? '1' : '0'}
            }}
        >
            <EmailForm visible={user === ""}
                       onSetUser={user => setUser(user)}
                       apikey={props.apikey}
                       httpServerPath={props.httpServerPath}
            />

            {user !== "" && <ChatEngine user={user}
                                        apikey={props.apikey}
                                        httpServerPath={props.httpServerPath}
                                        wsServerPath={props.wsServerPath}
            />}
        </div>

    )
}