# Oasis Voice Network
![Octavian Tails On The Phone](images/otter_phone.jpeg)

## Architecture

The Oasis Voice Network consists of Kamailio and Asterisk

![Network Diagram](images/Oasis%20Voice%20Network.png)

### Kamailio
* receives and sends WebRTC calls directly to the browser via SIP over WebSockets
* authenticates WebRTC users using the ephemeral auth module of Kamailio
* supports digest auth for making outbound pstn calls
* Identifies the carrier used by ingress calls


### Asterisk
* media anchors all calls
* transcodes between webrtc & classic media
* records calls (Comming Soon)


## Building
* Kamailio currently can build to AWS or Docker
* Asterions currently builds to AWS, docker is being investigated
* look at the asterisk and kamailio sub directories for more building instructions
