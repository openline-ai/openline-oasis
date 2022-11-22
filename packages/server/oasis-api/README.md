## Development

To run this service to run on your laptop you need the following environemnt vars to be set

| Variable                 | Meaning                                                                                                  |
|--------------------------|----------------------------------------------------------------------------------------------------------|
| MESSAGE_STORE_URL        | url of the GRPC interface of the message store                                                           |
| CHANNELS_API_URL         | url of the GPRC interface of the Channels API                                                            |
| CORS_URL                 | url of the frontend, needed to allow cros-site scripting                                                 |
| OASIS_API_SERVER_ADDRESS | interface and port to bind to, normally should be ":8006"                                                |
| WEBRTC_AUTH_SECRET       | Shared secret used for Ephemeral Authentication, should match AUTH_SECRET in your kamailio configuration |
| WEBRTC_AUTH_TTL          | Validity time in seconds of the Ephemeral Auth credentials                                               |
| OASIS_GRPC_PORT          | The grpc port that oasis-api uses                                                                        |
