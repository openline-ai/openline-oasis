import type { NextPage } from 'next'
import * as React from "react";
import { useEffect, useRef, useState } from "react";
import { Button } from "primereact/button";
import { useRouter } from "next/router";
import useWebSocket from 'react-use-websocket';
import axios from "axios";
import { InputTextarea } from "primereact/inputtextarea";
import { loggedInOrRedirectToLogin } from "../../utils/logged-in";
import { getSession, signOut, useSession } from "next-auth/react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
    faArrowRightFromBracket,
    faCaretDown,
    faPhone,
    faPhoneSlash,
    faRightLeft,
    faUserSecret,
    faMicrophone,
    faMicrophoneSlash,
    faPause,
    faPlay,
    faSquareXmark,
    faXmarkSquare
} from "@fortawesome/free-solid-svg-icons";

import { OverlayPanel } from "primereact/overlaypanel";
import { Menu } from "primereact/menu";
import { InputText } from "primereact/inputtext";
import Chat from "./chat";
import Moment from "react-moment";
import WebRTC from "../../components/webrtc/WebRTC";
import { FeedItem } from "../../model/feed-item";
import { ToastContainer } from "react-toastify";


const FeedPage: NextPage = () => {
    const router = useRouter()
    const { id } = router.query;

    const [feeds, setFeeds] = useState([] as FeedItem[]);
    const [selectedFeed, setSelectedFeed] = useState(id as string);

    const { lastMessage } = useWebSocket(`${process.env.NEXT_PUBLIC_WEBSOCKET_PATH}`, {
        onOpen: () => console.log('Websocket opened'),
        //Will attempt to reconnect on all close events, such as server shutting down
        shouldReconnect: (closeEvent) => true,
    });

    useEffect(() => {
        loadFeed();
    }, [id]);

    useEffect(() => {
        if (lastMessage && Object.keys(lastMessage).length !== 0) {
            loadFeed();
        }

    }, [lastMessage]);

    const loadFeed = function () {
        console.log("Reloading feed!");
        axios.get(`/oasis-api/feed`)
            .then(res => {
                setFeeds(res.data?.feedItems ?? []);
                if (!selectedFeed && res.data && res.data.feedItems && res.data.feedItems[0]) {
                    setSelectedFeed(res.data.feedItems[0].id);
                    router.push(`/feed?id=${res.data.feedItems[0].id}`, undefined, { shallow: true });
                }
            });
    }

    const { data: session, status } = useSession();
    const userSettingsContainerRef = useRef<OverlayPanel>(null);

    let userItems = [
        {
            label: 'My profile',
            icon: <FontAwesomeIcon icon={faUserSecret} className="mr-2" />,
            command: () => {
                router.push('/');
            }
        },
        {
            label: 'Logout',
            icon: <FontAwesomeIcon icon={faArrowRightFromBracket} className="mr-2" />,
            command: () => {
                signOut();
            }
        }
    ];

    // region WebRTC
    const phoneContainerRef = useRef<OverlayPanel>(null);
    const webrtc: React.RefObject<WebRTC> = useRef<WebRTC>(null);
    useEffect(() => {

        const refreshCredentials = () => {
            axios.get(`/oasis-api/call_credentials?service=sip&username=` + session?.user?.email)
                .then(res => {
                    console.error("Got a key: " + JSON.stringify(res.data));

                    webrtc.current?.setCredentials(res.data.username, res.data.password,
                        () => {
                            if (!webrtc.current?._ua) {
                                webrtc.current?.startUA()
                            }
                        });
                    setTimeout(() => {
                        refreshCredentials()
                    }, (res.data.ttl * 3000) / 4);
                });
        }
        if (session?.user?.email) {
            refreshCredentials();
        }
    }, [session?.user?.email]);

    const [callFrom, setCallFrom] = useState('');
    const [inCall, setInCall] = useState(false);
    const [showRefer, setShowRefer] = useState(false);
    const [inRefer, setInRefer] = useState(false);
    const [onHold, setOnHold] = useState(false);
    const [onMute, setOnMute] = useState(false);
    const [transferDest, setTransferDest] = useState('');


    const handleCall = (feedInitiator: any) => {
        let user = '';
        if (feedInitiator.phoneNumber) {
            user = feedInitiator.phoneNumber + "@oasis.openline.ai";
        } else {
            user = feedInitiator.email;
            const regex = /.*<(.*)>/;
            const matches = user.match(regex);
            if (matches) {
                user = matches[1];
            }
        }
        webrtc.current?.makeCall("sip:" + user);
    }
    const hangupCall = () => {
        webrtc.current?.hangupCall();

    }

    const showTransfer = () => {
        webrtc.current?.showTransfer();
    }

    const toggleMute = () => {
        if (onMute) {
            webrtc.current?.unMuteCall();
            setOnMute(false);
        } else {
            webrtc.current?.muteCall();
            setOnMute(true);
        }
    }

    const toggleHold = () => {
        if (onHold) {
            webrtc.current?.unHoldCall();
            setOnHold(false);
        } else {
            webrtc.current?.holdCall();
            setOnHold(true);
        }
    }

    //endregion
    const makeButton = (number: string) => {
        return <button className="btn btn-primary btn-lg m-1" key={number}
            onClick={() => { webrtc.current?.sendDtmf(number) }}>{number}</button>
    }

    let dialpad_matrix = new Array(4)
    for (let i = 0, digit = 1; i < 3; i++) {
        dialpad_matrix[i] = new Array(3);
        for (let j = 0; j < 3; j++, digit++) {
            dialpad_matrix[i][j] = makeButton(digit.toString())
        }
    }
    dialpad_matrix[3] = new Array(3);
    dialpad_matrix[3][0] = makeButton("*")
    dialpad_matrix[3][1] = makeButton("0")
    dialpad_matrix[3][2] = makeButton("#")

    let dialpad_rows = []
    for (let i = 0; i < 4; i++) {
        dialpad_rows.push(<div className="d-flex flex-row justify-content-center">{dialpad_matrix[i]}</div>)
    }

    function setTopbarColor(newColor: any) {
        document.documentElement.style.setProperty('--topbar-background', newColor);
    }

    useEffect(() => {
        if (inCall) {
            setTopbarColor('#FFCCCB');
        } else {
            setTopbarColor('#FFFFFF');
        }
    }, [inCall]);

    return (
        <div className="flex w-full h-full">
            <ToastContainer position="top-center"
                autoClose={3000}
                closeOnClick={true}
                hideProgressBar={true}
                theme="colored" />
            {
                process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL &&
                <WebRTC
                    ref={webrtc}
                    websocket={`${process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL}`}
                    from={"sip:" + session?.user?.email}
                    notifyCallFrom={(callFrom: string) => setCallFrom(callFrom)}
                    updateCallState={(state: boolean) => setInCall(state)}
                    autoStart={false}
                />
            }

            <div className="flex flex-column flex-grow-0 h-full overflow-hidden"
                style={{ width: '350px', background: 'white', borderRight: '1px rgb(235, 235, 235) solid' }}>

                <div className='flex p-3'>
                    <InputText placeholder={'Search'} className='w-full' />
                </div>

                    <div className='flex flex-column pl-3 pr-3 mb-3 overflow-x-hidden overflow-y-auto'>
                        {
                            feeds.map((f: FeedItem) => {
                                let className = 'flex w-full align-content-center align-items-center p-3 mb-2 contact-hover';
                                if (selectedFeed === f.id) {
                                    className += ' selected'
                                }

                                var t = new Date(1970, 0, 1);
                                t.setSeconds(f.lastTimestamp.seconds);

                                return <div key={f.id} className={className} onClick={() => {
                                    setSelectedFeed(f.id);
                                    //change the URL to allow a bookmark
                                    router.push(`/feed?id=${f.id}`, undefined, { shallow: true });
                                }
                                }>
                                    <div className='flex flex-column flex-grow-1 mr-3' style={{ minWidth: '0' }}>
                                        <div className='mb-2'>
                                            {
                                                f.initiatorFirstName &&
                                                f.initiatorFirstName + ' ' + f.initiatorLastName}
                                            {
                                                !f.initiatorFirstName &&
                                                f.initiatorUsername
                                            }
                                        </div>
                                        <div className='text-500' style={{
                                            fontSize: '12px',
                                            textOverflow: 'ellipsis',
                                            whiteSpace: "nowrap",
                                            overflow: "hidden"
                                        }}>
                                            {f.lastContentPreview}
                                        </div>
                                    </div>

                                    <div className='flex flex-column text-right'>
                                        <Moment className="text-sm text-gray-600" date={t}
                                            format={'MMM D, YYYY'}></Moment>
                                        <Moment className="text-sm text-gray-600" date={t}
                                            format={'HH:mma'}></Moment>
                                    </div>
                                </div>
                            })
                        }
                    </div>

            </div>

                <div className='flex h-full w-full flex-column'>
                    <div className='openline-top-bar'>
                        <div className="flex align-items-center justify-content-end">

                            {
                                !inCall &&
                                <>
                                    <Button className="p-button-rounded p-button-success p-2"
                                        onClick={(e: any) => phoneContainerRef?.current?.toggle(e)}>
                                        <FontAwesomeIcon icon={faPhone} fontSize={'16px'} />
                                    </Button>

                                    <OverlayPanel ref={phoneContainerRef} dismissable>
                                        <div style={{position: "relative", width: "100%", height: "100%"}}>
                                        <div className='font-bold text-center'>In call with</div>
                                        <div className='font-bold text-center mb-3'>{dialpad_rows}</div>

                                        <div className='font-bold text-center mb-3'>{callFrom}</div>
                                        <div className='mb-3'>
                                        <Button onClick={() => toggleMute()} className="mr-2">
                                            <FontAwesomeIcon icon={onMute ? faMicrophone : faMicrophoneSlash} className="mr-2" /> {onMute ? "Unmute" : "Mute"}
                                        </Button>
                                        <Button onClick={() => toggleHold()} className="mr-2">
                                            <FontAwesomeIcon icon={onHold ? faPlay : faPause} className="mr-2" /> {onHold ? "Release hold" : "Hold"}
                                        </Button>
                                        <Button onClick={() => hangupCall()} className='p-button-danger mr-2'>
                                            <FontAwesomeIcon icon={faPhoneSlash} className="mr-2" /> Hangup
                                        </Button>
                                        <Button onClick={() => setShowRefer(!showRefer)} className='p-button-success mr-2'>
                                            <FontAwesomeIcon icon={showRefer?faXmarkSquare:faRightLeft} className="mr-2" /> Transfer
                                        </Button>
                                        </div>
                                        {showRefer &&
                                        <>
                                        <div className="w-full text-center align-items-center mb-3">
                                            <InputTextarea className="mr-2"
                                                value={transferDest}
                                                onChange={(e) => setTransferDest(e.target.value)}
                                                autoResize
                                                rows={1}
                                                placeholder="Transfer to"
                                                onKeyPress={(e) => {
                                                    if (e.shiftKey && e.key === "Enter") {
                                                        return true
                                                    }
                                                    if (e.key === "Enter") {
                                                        e.preventDefault();
                                                    }
                                                }}
                                                style={{
                                                    borderColor: "black", //Do not set as none!! It breaks InputTextarea autoResize
                                                    boxShadow: "none"
                                                }}
                                            />
                                            <span className="h-full align-items-top">
                                            <Button onClick={() => setInRefer(!inRefer)} className='p-button-success h-full mr-2'>
                                                <FontAwesomeIcon icon={faRightLeft} className="mr-2" />
                                            </Button>
                                            </span>
                                        </div>
                                        </>}
                                        {inRefer && <div 
                                            style={{
                                                position: "absolute",
                                                zIndex: 2000,
                                                width: "100%",
                                                height: "100%",
                                                top: "0%",
                                                background: "#FFFFFFFF",
                                            }}>
                                            <div
                                            style={{  margin: 0,
                                                position: "absolute",
                                                top: "50%",
                                                transform: "translateY(-50%)",
                                                width: "100%"
                                                }}>
                                                    <div className="w-full text-center align-items-center mb-3">
                                                    Transfering call to: {transferDest}
                                                    </div>
                                                    <div className="w-full text-center align-items-center mb-3">
                                                    <FontAwesomeIcon icon={faRightLeft} className="mr-2" />
                                                    </div>
                                            </div>

                                            </div>

                                            
                                        }
                                    </div>
                                    </OverlayPanel>
                                </>
                            }

                            <Button className="flex-none px-3 m-3"
                                onClick={(e: any) => userSettingsContainerRef?.current?.toggle(e)}>
                                <FontAwesomeIcon icon={faUserSecret} className="mr-2" />
                                <span className='flex-grow-1'>{session?.user?.email}</span> {/* TODO: Add name */}
                                <FontAwesomeIcon icon={faCaretDown} className="ml-2" />
                            </Button>

                            <OverlayPanel ref={userSettingsContainerRef} dismissable>
                                <Menu model={userItems} style={{ border: 'none' }} />
                            </OverlayPanel>

                        </div>
                    </div>
                </div>
                {
                    selectedFeed &&
                    <Chat
                        feedId={selectedFeed}
                        inCall={inCall}
                        handleCall={(feedInitiator: any) => handleCall(feedInitiator)}
                        hangupCall={() => hangupCall()}
                        showTransfer={() => showTransfer()}
                    />
                }
            </div>


        </div>
    );
}

export async function getServerSideProps(context: any) {
    return loggedInOrRedirectToLogin(await getSession(context));
}

export default FeedPage
