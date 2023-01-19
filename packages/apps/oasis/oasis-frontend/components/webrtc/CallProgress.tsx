import type { NextPage } from 'next'
import * as React from "react";
import { useEffect, useRef, useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
    faPhone,
    faPhoneSlash,
    faRightLeft,
    faMicrophone,
    faMicrophoneSlash,
    faPause,
    faPlay,
    faXmarkSquare
} from "@fortawesome/free-solid-svg-icons";

import { OverlayPanel } from "primereact/overlaypanel";
import { Button } from "primereact/button";
import WebRTC from "../../components/webrtc/WebRTC";
import { InputTextarea } from "primereact/inputtextarea";
import SuggestionList from '../SuggestionList';

interface CallProgressProps {
    inCall: boolean
    webrtc: React.RefObject<WebRTC>
    callFrom: string
    referStatus: string
    getContactSuggestions: Function
}


const CallProgress: NextPage<CallProgressProps> = (props:CallProgressProps) => {
    const [transferDest, setTransferDest] = useState('');
    const phoneContainerRef = useRef<OverlayPanel>(null);
    const [showRefer, setShowRefer] = useState(false);
    const [inRefer, setInRefer] = useState(false);
    const [onHold, setOnHold] = useState(false);
    const [onMute, setOnMute] = useState(false);
    const [referProgressString, setReferProgressString] = useState('');

    const toggleMute = () => {
        if (onMute) {
            props.webrtc.current?.unMuteCall();
            setOnMute(false);
        } else {
            props.webrtc.current?.muteCall();
            setOnMute(true);
        }
    }

    const toggleHold = () => {
        if (onHold) {
            props.webrtc.current?.unHoldCall();
            setOnHold(false);
        } else {
            props.webrtc.current?.holdCall();
            setOnHold(true);
        }
    }

    const transferCall = () => {
        console.log("transferCall")

        setReferProgressString((''))
        setInRefer(true);
        props.webrtc.current?.transferCall(transferDest);
    }

    const clearRefer = () => {
        setReferProgressString('');
        setInRefer(false);
        props.webrtc.current?.hangupCall();
    }

    useEffect(() => {
        if (props.referStatus === 'referProgress') {
            setReferProgressString('Refer in progress');
        } else if (props.referStatus === 'referSuccess') {
            setReferProgressString('Refer successful');
        } else if (props.referStatus === 'referFailed') {
            setReferProgressString('Refer failed');
        }
        switch(props.referStatus) {
            case 'requestSucceeded':
                setReferProgressString('Transferring Call....')
                break;
            case 'trying':
                setReferProgressString('Trying to transfer call....')
                break;
            case 'progress':
                setReferProgressString('Ringing....')
                break;
            case 'requestFailed':
                setReferProgressString('Failed to contact remote party')
                setTimeout(clearRefer, 2000);
                break;
            case 'failed':
                setReferProgressString('Transfer Request rejected')
                setTimeout(clearRefer, 2000);
                break;
            case 'accepted':
                setReferProgressString('Call Transferred')
                setTimeout(clearRefer, 2000);
                break;

            default:
        }


    },[props.referStatus])


    const makeButton = (number: string) => {
        return <button className="btn btn-primary btn-lg m-1" key={"dtmf-"+number}
            onClick={() => { props.webrtc.current?.sendDtmf(number) }}>{number}</button>
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
        dialpad_rows.push(<div key={"dtmf-row-" + i} className="d-flex flex-row justify-content-center">{dialpad_matrix[i]}</div>)
    }

    return (
        <>
        {
            props.inCall &&
            <>
                <Button className="p-button-rounded p-button-success p-2"
                    onClick={(e: any) => phoneContainerRef?.current?.toggle(e)}>
                    <FontAwesomeIcon icon={faPhone} fontSize={'16px'} />
                </Button>

                <OverlayPanel ref={phoneContainerRef} dismissable>
                    <div style={{position: "relative", width: "100%", height: "100%"}}>
                    <div className='font-bold text-center'>In call with</div>
                    <div className='font-bold text-center mb-3'>{dialpad_rows}</div>

                    <div className='font-bold text-center mb-3'>{props.callFrom}</div>
                    <div className='mb-3'>
                    <Button onClick={() => toggleMute()} className="mr-2">
                        <FontAwesomeIcon icon={onMute ? faMicrophone : faMicrophoneSlash} className="mr-2" /> {onMute ? "Unmute" : "Mute"}
                    </Button>
                    <Button onClick={() => toggleHold()} className="mr-2">
                        <FontAwesomeIcon icon={onHold ? faPlay : faPause} className="mr-2" /> {onHold ? "Release hold" : "Hold"}
                    </Button>
                    <Button onClick={() => props.webrtc.current?.hangupCall()} className='p-button-danger mr-2'>
                        <FontAwesomeIcon icon={faPhoneSlash} className="mr-2" /> Hangup
                    </Button>
                    <Button onClick={() => setShowRefer(!showRefer)} className='p-button-success mr-2'>
                        <FontAwesomeIcon icon={showRefer?faXmarkSquare:faRightLeft} className="mr-2" /> Transfer
                    </Button>
                    </div>
                    {showRefer &&
                    <>
                    <div>
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
                        <Button onClick={transferCall} className='p-button-success h-full mr-2'>
                            <FontAwesomeIcon icon={faRightLeft} className="mr-2" />
                        </Button>
                        </span>
                    </div>
                    <div>
                        <SuggestionList currentValue={transferDest} getSuggestions={props.getContactSuggestions} setCurrentValue={setTransferDest}></SuggestionList>
                    </div>
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
                                <div key="referProgress" className="w-full text-center align-items-center mb-3">
                                {referProgressString}
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
        </>

    );
}

export default CallProgress
