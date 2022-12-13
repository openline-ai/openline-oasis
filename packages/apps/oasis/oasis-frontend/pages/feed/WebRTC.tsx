import * as React from 'react'
import * as JsSIP from 'jssip';

import {Button} from "primereact/button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPhone, faPhoneSlash} from "@fortawesome/free-solid-svg-icons";
import {InputText} from "primereact/inputtext";

import {
    EndEvent,
    IncomingAckEvent,
    IncomingEvent,
    OutgoingAckEvent,
    OutgoingEvent,
    RTCSessionEventMap,
    RTCSession,
    ReferOptions,
} from "jssip/lib/RTCSession";
import {IncomingRTCSessionEvent, OutgoingRTCSessionEvent, UAConfiguration} from "jssip/lib/UA";

interface WebRTCState {
    inCall: boolean
    websocket: string
    from: string
    updateCallState: Function
    callerId: string
    ringing: boolean
    autoStart: boolean
    username?: string
    password?: string
    transferDestination: string
    refer: boolean
    referStatus: string

}

interface WebRTCProps {
    websocket: string
    from: string
    updateCallState: Function
    autoStart?: boolean
}

export default class WebRTC extends React.Component<WebRTCProps> {
    state: WebRTCState
    _ua: JsSIP.UA | null
    _session:  RTCSession | null | undefined
    remoteVideo: React.RefObject<HTMLVideoElement>

    //setState:Function

    constructor(props: WebRTCProps) {
        super(props);
        this.state =
            {
                inCall: false,
                websocket: props.websocket,
                from: props.from,
                updateCallState: props.updateCallState,
                callerId: "",
                ringing: false,
                autoStart: false,
                transferDestination: "",
                refer: false,
                referStatus: ""
            };

        if (props.autoStart) {
            this.state.autoStart = props.autoStart;
        }

        this._ua = null;
        this.remoteVideo = React.createRef();
        //this.setState({inCall: this.state.inCall, callerId: this.state.callerId, ringing: this.state.ringing});
        this._session = null;

    }

    answerCall() {
        this.setState({inCall: true, ringing: false});
        this.state.updateCallState(true);
        this._session?.answer();


    }

    showTransfer() {
        this.setState({refer: !this.state.refer});

    }

    transferCall() {
        let transferDest = this.state.transferDestination;
        var localScope = this;
        localScope.setState({referStatus: ''})

        let eventHandlers = {
            'requestSucceeded': function (e: any) {
                console.log('xfer is accepted');
                localScope.setState({referStatus: 'Transferring Call....'})

            },
            'requestFailed': function (e: any) {
                console.log('Faled to contact remote party cause: ' + JSON.stringify(e.cause));
                localScope.setState({referStatus: 'Faled to contact remote party cause: ' + JSON.stringify(e.cause)})

            },
            'failed': function (e: any) {
                console.log('Transfer Request rejected with cause: ' + JSON.stringify(e.cause));
                localScope.setState({referStatus: 'Faled to contact remote party cause: ' + JSON.stringify(e.cause)})

            },
            'accepted': function (e: IncomingAckEvent | OutgoingAckEvent) {
                console.log('call confirmed');
                localScope.setState({referStatus: ''})
                localScope.setState({refer: false});

            }
        };
        let options : ReferOptions  = {
            'eventHandlers': eventHandlers,
        }
        if (transferDest.indexOf('@') === -1) {
            transferDest = transferDest + "@agent.openline.ai";
        }
        if (!transferDest.startsWith("sip:")) {
            transferDest = "sip:" + transferDest;
        }
        this._session?.refer(transferDest, options);


    }

    hangupCall() {
        this.setState({inCall: false, ringing: false});
        this.state.updateCallState(false);
        if (this._session) {
            this._session.terminate();
        }
    }

    setCredentials(user: string, pass: string, callback?: (() => void)) {
        if (!callback) {
            this.setState({username: user, password: pass});
        } else {
            this.setState({username: user, password: pass}, callback);
        }
    }

    makeCall(destination: string) {
        var localScope = this;
        this.setState({inCall: true, ringing: false});
        let eventHandlers: Partial<RTCSessionEventMap> = {
            'progress': function (e: IncomingEvent | OutgoingEvent) {
                console.log('call is in progress');
                localScope.setState({inCall: true});
                localScope.state.updateCallState(true);
            },
            'failed': function (e: EndEvent) {
                console.log('call failed with cause: ' + JSON.stringify(e.cause));
                localScope.setState({inCall: false})
                localScope.state.updateCallState(false);
            },
            'ended': function (e: EndEvent) {
                console.log('call ended with cause: ' + JSON.stringify(e.cause));
                localScope.setState({inCall: false});
                localScope.state.updateCallState(false);
            },
            'confirmed': function (e: IncomingAckEvent | OutgoingAckEvent) {
                console.log('call confirmed');
                localScope.setState({inCall: true});
                localScope.state.updateCallState(true);
            }
        };

        var options: any = {
            'eventHandlers': eventHandlers,
            'mediaConstraints': {'audio': true, 'video': true},
        };
        if (process.env.NEXT_PUBLIC_TURN_SERVER)
            options['pcConfig'] = {
                'iceServers': [
                    {
                        'urls': [process.env.NEXT_PUBLIC_TURN_SERVER],
                        'username': process.env.NEXT_PUBLIC_TURN_USER,
                        'credential': process.env.NEXT_PUBLIC_TURN_USER
                    },
                ]
            };

        this.setState({inCall: true});
        this._session = this._ua?.call(destination, options);
        var peerconnection = this._session?.connection;
        peerconnection?.addEventListener('addstream', (event: any) => {
                if (this.remoteVideo.current) {
                    this.remoteVideo.current.srcObject = event.stream;
                }
                this.remoteVideo.current?.play();
            }
        )
    }


    componentDidMount() {
        if (this.state.autoStart) {
            this.startUA();
        }
    }

    stopUA() {
        if (!this._ua) {
            console.log("UA not yet started! ignoring request");
            return;
        }

        this._ua.stop();
        this._ua = null;
    }

    startUA() {
        if (this._ua) {
            console.log("UA already started! ignoring request");
            return;
        }
        let socket: JsSIP.Socket = new JsSIP.WebSocketInterface(this.state.websocket);
        let configuration: UAConfiguration = {
            sockets: [socket],
            uri: this.state.from
        };

        if (this.state.username) {
            configuration.authorization_user = this.state.username;
            configuration.password = this.state.password;
        }

        console.error("Got a configuration: " + JSON.stringify(configuration));
        JsSIP.debug.enable('JsSIP:*');
        this._ua = new JsSIP.UA(configuration);
        this._ua.on('newRTCSession', ({
                                          originator,
                                          session: rtcSession,
                                          request
                                      }: IncomingRTCSessionEvent | OutgoingRTCSessionEvent) => {
            if (originator === 'local')
                return;

            this._session = rtcSession;
            this.setState({ringing: true});
            this.setState({inCall: true});
            this.setState({callerId: rtcSession.remote_identity.uri.toString()});
            this.state.updateCallState(true);
            console.error("Got a call for " + rtcSession.remote_identity.uri.toString());
            rtcSession.on('accepted', () => {
                    if (this.remoteVideo.current) {
                        this.remoteVideo.current.srcObject = (this._session?.connection.getRemoteStreams()[0]?this._session?.connection.getRemoteStreams()[0]: null);
                        this.remoteVideo.current.play();
                    }
                }
            );
            rtcSession.on('ended', () => {
                this.setState({inCall: false, ringing: false});
                this.state.updateCallState(false);
            });

        });
        this._ua.start();
    }

    render() {

        return (
            <>
                <div style={{position: "absolute", zIndex: 9, width: 'calc(100% - 150px)'}} hidden={!this.state.inCall}>
                    <video controls={false} hidden={!this.state.inCall}
                           ref={this.remoteVideo} autoPlay>

                    </video>
                    <div style={{
                        position: "absolute",
                        top: "50%",
                        left: "50%",
                        width: "50%",
                        textAlign: "center",
                        transform: "translate(-50%, -50%)",
                        background: "lightgrey",
                        borderRadius: '3px',
                        border: "1px solid black"
                    }} hidden={!this.state.inCall || !this.state.ringing}>
                        Incomming call from {this.state.callerId}<br/>
                        <Button onClick={() => this.answerCall()} className='p-button-text'>
                            <FontAwesomeIcon icon={faPhone} style={{color: 'black'}}/>
                        </Button>
                        <Button onClick={() => this.hangupCall()} className='p-button-text'>
                            <FontAwesomeIcon icon={faPhoneSlash} style={{color: 'black'}}/>
                        </Button>
                    </div>
                    <div style={{
                        position: "absolute",
                        top: "50%",
                        left: "50%",
                        width: "50%",
                        textAlign: "center",
                        transform: "translate(-50%, -50%)",
                        background: "lightgrey",
                        borderRadius: '3px',
                        border: "1px solid black"
                    }} hidden={!this.state.refer}>
                        Specify desitnation for Call Transfer<br/>
                        <div>{this.state.referStatus}</div>
                        <InputText style={{width: 'calc(100% - 150px)'}} value={this.state.transferDestination}
                               onChange={(e) => this.setState({transferDestination:e.target.value})}/>
                        <Button onClick={() => this.transferCall()} className='p-button-text'>
                            <FontAwesomeIcon icon={faPhone} style={{color: 'black'}}/>
                        </Button>

                    </div>
                </div>
            </>

        )
    }
}
