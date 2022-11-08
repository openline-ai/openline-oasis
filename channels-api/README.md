

### generate
https://grpc.io/docs/protoc-installation/
https://grpc.io/docs/languages/go/quickstart/

To build the project run

```
make
```

## Development

To run this service to run on your laptop you need the following environemnt vars to be set

| Variable                    | Meaning                                                      |
|-----------------------------|--------------------------------------------------------------|
| MESSAGE_STORE_URL           | url of the GRPC interface of the message store               |
| CHANNELS_API_SERVER_ADDRESS | ip:port to bind for the rest api, normally should be ":8013" |
| GRPC_PORT                   | port used for the channel-api grpc interface, should be 9013 |
| SMTP_SERVER_ADDRESS         | hostname of smtp server to connect to                        |
| SMTP_SERVER_USER            | user to authenticate with the smtp server as                 |
| SMTP_SERVER_PASSWORD        | password to authenticate with the smtp server                |
| SMTP_FROM_USER              | email address to send email as                               |