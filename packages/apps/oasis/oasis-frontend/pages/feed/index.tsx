import type {NextPage} from 'next'
import {useEffect, useRef, useState} from "react";
import {Button} from "primereact/button";
import {useRouter} from "next/router";
import useWebSocket from 'react-use-websocket';
import axios from "axios";
import {loggedInOrRedirectToLogin} from "../../utils/logged-in";
import {getSession, signOut, useSession} from "next-auth/react";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faArrowRightFromBracket, faCaretDown, faUserSecret} from "@fortawesome/free-solid-svg-icons";
import {OverlayPanel} from "primereact/overlaypanel";
import {Menu} from "primereact/menu";
import {InputText} from "primereact/inputtext";
import Chat from "./chat";
import Moment from "react-moment";


const FeedPage: NextPage = () => {
    const router = useRouter()
    const {id} = router.query;

    const [feeds, setFeeds] = useState([] as any)
    const [selectedFeed, setSelectedFeed] = useState(id);

    const {lastMessage} = useWebSocket(`${process.env.NEXT_PUBLIC_WEBSOCKET_PATH}`, {
        onOpen: () => console.log('Websocket opened'),
        //Will attempt to reconnect on all close events, such as server shutting down
        shouldReconnect: (closeEvent) => true,
    });

    useEffect(() => {
        axios.get(`/server/feed`)
                .then(res => {
                    res.data?.feedItems.forEach((f: any) => {
                        f.updatedOn.dateTime = toDateTime(f.updatedOn.seconds);
                    });
                    setFeeds(res.data?.feedItems)
                    if (!selectedFeed) {
                        setSelectedFeed(res.data.feedItems[0].id);
                        router.push(`/feed?id=${res.data.feedItems[0].id}`, undefined, {shallow: true});
                    }
                })
    }, []);

    useEffect(() => {
        if (lastMessage && Object.keys(lastMessage).length !== 0) {
            handleWebsocketMessage(lastMessage);
        }

    }, [lastMessage, setFeeds]);

    const handleWebsocketMessage = function (msg: any) {
        console.log("Got a new feed!");
        axios.get(`/server/feed`)
                .then(res => {
                    res.data?.feedItems.forEach((f: any) => {
                        f.updatedOn.dateTime = toDateTime(f.updatedOn.seconds);
                    });
                    setFeeds(res.data?.feedItems);
                    if (!selectedFeed) {
                        setSelectedFeed(res.data.feedItems[0].id);
                        router.push(`/feed?id=${res.data.feedItems[0].id}`, undefined, {shallow: true});
                    }
                });
    }

    const {data: session, status} = useSession();
    const userSettingsContainerRef = useRef<OverlayPanel>(null);

    let userItems = [
        {
            label: 'My profile',
            icon: <FontAwesomeIcon icon={faUserSecret} className="mr-2"/>,
            command: () => {
                router.push('/');
            }
        },
        {
            label: 'Logout',
            icon: <FontAwesomeIcon icon={faArrowRightFromBracket} className="mr-2"/>,
            command: () => {
                signOut();
            }
        }
    ];

    const toDateTime = function (secs: any) {
        var t = new Date(1970, 0, 1);
        t.setSeconds(secs);
        return t;
    }

    return (
            <>
                <div className="flex w-full h-full">

                    <div className="flex flex-column flex-grow-0 h-full overflow-hidden"
                         style={{width: '30%', background: 'white'}}>

                        <div className="flex flex-row align-items-center justify-content-between pt-3 pl-3 pr-3">

                            <Button className="dark-button flex-grow-1"
                                    onClick={(e: any) => userSettingsContainerRef?.current?.toggle(e)}>
                                <FontAwesomeIcon icon={faUserSecret} className="mr-2"/>
                                <span className='flex-grow-1'>{session?.user?.email}</span>
                                <FontAwesomeIcon icon={faCaretDown} className="ml-2"/>
                            </Button>

                            <OverlayPanel ref={userSettingsContainerRef} dismissable>
                                <Menu model={userItems} style={{border: 'none'}}/>
                            </OverlayPanel>

                        </div>

                        <div className='flex p-3'>
                            <InputText placeholder={'Search'} className='w-full'/>
                        </div>

                        <div className='flex flex-column pl-3 pr-3 mb-3 overflow-x-hidden overflow-y-auto'>
                            {
                                feeds.map((f: any) => {
                                    let className = 'flex w-full align-content-center align-items-center p-3 mb-2 contact-hover';
                                    if ( selectedFeed === f.id ) {
                                        className += ' selected'
                                    }
                                    return <div key={f.email} className={className} onClick={() => {
                                        setSelectedFeed(f.id);
                                        //change the URL to allow a bookmark
                                        router.push(`/feed?id=${f.id}`, undefined, {shallow: true});
                                    }
                                    }>
                                        {/*<div style={{height: "10px", width: "18px", borderRadius: "100px", background: "#7626FA"}}></div>*/}

                                        <div className='flex flex-column flex-grow-1 mr-3' style={{minWidth: '0'}}>
                                            <div className='mb-2'>
                                                {
                                                        f.contactFirstName &&
                                                        f.contactFirstName + ' ' + f.contactLastName}
                                                {
                                                        !f.contactFirstName &&
                                                        f.contactEmail
                                                }
                                            </div>
                                            <div className='text-500' style={{
                                                fontSize: '12px',
                                                textOverflow: 'ellipsis',
                                                whiteSpace: "nowrap",
                                                overflow: "hidden"
                                            }}>
                                                {f.message}
                                            </div>
                                        </div>

                                        <div className='flex flex-column'>
                                            <Moment className="text-sm text-gray-600" date={f.updatedOn.dateTime}
                                                    format={'d.MM.yy'}></Moment>
                                            <Moment className="text-sm text-gray-600" date={f.updatedOn.dateTime}
                                                    format={'HH:mm'}></Moment>
                                        </div>
                                    </div>
                                })
                            }
                        </div>

                    </div>

                    <div className='flex flex-grow-1 w-full'>

                        {
                                selectedFeed &&
                                <Chat id={selectedFeed}/>
                        }

                    </div>


                </div>

            </>
    );
}

export async function getServerSideProps(context: any) {
    return loggedInOrRedirectToLogin(await getSession(context));
}

export default FeedPage
