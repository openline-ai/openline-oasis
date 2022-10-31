import  {useEffect, useState} from 'react'
import * as React from 'react'
import {Client, IFrame, IMessage} from '@stomp/stompjs'

const client = new Client();
let reactMessage:IFrame|null = null;
let reactSetMessage:Function = (message:string) => {}

export const useStomp = () => {
    const [message, setMessage] = useState<IFrame>({command: "", headers: {}, binaryBody: new Uint8Array(), isBinaryBody: false, body: "" });

    reactMessage = message;
    reactSetMessage = setMessage;
    return message;
}

export const configureStomp = (url:string, topic:string) => {
    client.configure({
        brokerURL: url,
        onConnect: () => {
            console.log('onConnect');
            client.subscribe(topic, message => {
                reactSetMessage(message);
            });
        },
        // Helps during debugging, remove in production
        debug: (str) => {
            console.log(new Date(), str);
        }
    });
    client.activate();
}

export const doNothing = () => {}

export default doNothing