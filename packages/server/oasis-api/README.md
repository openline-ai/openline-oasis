## Development

This service uses the environment variables described below. The env files have a default value if not provided ( check .env file )

| Variable                      | Meaning                                                                                                  |
|-------------------------------|----------------------------------------------------------------------------------------------------------|
| MESSAGE_STORE_URL             | url of the GRPC interface of the message store                                                           |
| MESSAGE_STORE_API_KEY         | message store API key                                                                                    |
| CHANNELS_API_URL              | url of the GPRC interface of the Channels API                                                            |
| CORS_URL                      | url of the frontend, needed to allow cros-site scripting                                                 |
| OASIS_API_SERVER_ADDRESS      | interface and port to bind to, normally should be ":8006"                                                |
| WEBRTC_AUTH_SECRET            | Shared secret used for Ephemeral Authentication, should match AUTH_SECRET in your kamailio configuration |
| WEBRTC_AUTH_TTL               | Validity time in seconds of the Ephemeral Auth credentials                                               |
| OASIS_GRPC_PORT               | The grpc port that oasis-api uses                                                                        |
| OASIS_API_KEY                 | API key the server expects to see in rest requests received                                              |
| WEBSOCKET_PING_INTERVAL       | Ping interval in seconds to monitor websocket connections                                                |
| POSTGRES_HOST                 | hostname/ip of the postgres db to connect to                                                             |
| POSTGRES_PORT                 | port of the database normally should be 5432                                                             |
| POSTGRES_USER                 | username to connect to the datbase with                                                                  |
| POSTGRES_DB                   | name of the postgres db to use                                                                           |
| POSTGRES_PASSWORD             | password to use to connect to the database                                                               |
| POSTGRES_DB_MAX_CONN          | (optional)                                                                                               |
| POSTGRES_DB_MAX_IDLE_CONN     | (optional)                                                                                               |     
| POSTGRES_DB_CONN_MAX_LIFETIME | (optional)                                                                                               | 
| NEO4J_TARGET                  | The target neo4j instance to connect to                                                                  |
| NEO4J_AUTH_USER               | The username to use for authentication in the neo4j database                                             |
| NEO4J_AUTH_PWD                | The password to use for authentication in the neo4j database                                             |
| NEO4J_AUTH_REALM              | (optional) The realm to use for authentication in the neo4j database                                     |
| NEO4J_MAX_CONN_POOL_SIZE      | (optional) The maximum number of connections to the neo4j database                                       |
| NEO4J_LOG_LEVEL               | (optional) The log level to use for the neo4j driver                                                     |