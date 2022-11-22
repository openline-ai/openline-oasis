import React, {useEffect, useRef, useState} from "react";
import Avatar from "../Avatar/Avatar";
import WebChatWindow from "../WebChatWindow/WebChatWindow";
import OpenlineTracker from "./OpenlineTracker";

interface WebChatProps {
    apikey: string
    httpServerPath: string
    wsServerPath: string
    trackerEnabled: boolean
    trackerAppId: string
    trackerId: string
    trackerCollectorUrl: string
    trackerBufferSize: string
    trackerMinimumVisitLength: string
    trackerHeartbeatDelay: string
}

export default function WebChat(props: WebChatProps) {
    const componentRef = useRef();

    const [visible, isVisible] = useState(false);

    useEffect(() => {
        document.addEventListener("click", handleClick);
        return () => document.removeEventListener("click", handleClick);

        function handleClick(e: any) {
            if (componentRef && componentRef.current) {
                const ref: any = componentRef.current
                if (!ref.contains(e.target)) {
                    isVisible(false)
                }
            }
        }
    }, []);

    return (
        <div ref={componentRef as any}>
            <WebChatWindow visible={visible}
                           apikey={props.apikey}
                           httpServerPath={props.httpServerPath}
                           wsServerPath={props.wsServerPath}
            />
            <Avatar
                onClick={() => {
                    isVisible(true)
                }}
                style={{
                    position: 'fixed',
                    bottom: '24px',
                    right: '24px',
                }}
            />
            <OpenlineTracker
                enabled={props.trackerEnabled}
                appId={props.trackerAppId}
                trackerId={props.trackerId}
                collectorUrl={props.trackerCollectorUrl}
                bufferSize={props.trackerBufferSize}
                minimumVisitLength={props.trackerMinimumVisitLength}
                heartbeatDelay={props.trackerHeartbeatDelay}
            />
        </div>
    )

}