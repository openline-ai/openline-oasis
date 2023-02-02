

### generate
https://grpc.io/docs/protoc-installation/
https://grpc.io/docs/languages/go/quickstart/

To build the project run

```
make
```

## Development

This service uses the environment variables described below. The env files have a default value if not provided ( check .env file )

| Variable                      | Meaning                                                                                |
|-------------------------------|----------------------------------------------------------------------------------------|
| MESSAGE_STORE_URL             | url of the GRPC interface of the message store                                         |
| MESSAGE_STORE_API_KEY         | message store API key                                                                  |
| CHANNELS_API_SERVER_ADDRESS   | ip:port to bind for the rest api, normally should be ":8013"                           |
| CHANNELS_GRPC_PORT            | port used for the channel-api grpc interface, should be 9013                           |
| MAIL_API_KEY                  | The api key used to validated received emails, must mach what is set in the AWS Lambda |
| OASIS_API_URL                 | IP & port of the GRPC interface of oasis api                                           |
| CHANNELS_API_CORS_URL         | url of the frontend, needed to allow cros-site scripting                               |
| WEBCHAT_API_KEY               | The api key used to validated received messages and login requests                     |
| WEBSOCKET_PING_INTERVAL       | Ping interval in seconds to monitor websocket connections                              |
| POSTGRES_HOST                 | hostname/ip of the postgres db to connect to                                           |
| POSTGRES_PORT                 | port of the database normally should be 5432                                           |
| POSTGRES_USER                 | username to connect to the datbase with                                                |
| POSTGRES_DB                   | name of the postgres db to use                                                         |
| POSTGRES_PASSWORD             | password to use to connect to the database                                             |
| POSTGRES_DB_MAX_CONN          | (optional)                                                                             |
| POSTGRES_DB_MAX_IDLE_CONN     | (optional)                                                                             |     
| POSTGRES_DB_CONN_MAX_LIFETIME | (optional)                                                                             | 


## Setting up gmail in local environment

follow the procedure in https://developers.google.com/gmail/api/quickstart/go
start ngrok to tunnel to channel-api
```
ngrok http 8013
```

* create a credential of type oauth client id
* select web application as application type
* add http://localhost:3006 as authorized javascript origin
* add https://(your ngrok url)/auth as authorized redirect uri
* set GMAIL_CLIENT_ID to the client id
* set GMAIL_CLIENT_SECRET to the client secret
* set GMAIL_REDIRECT_URIS to the redirect url


