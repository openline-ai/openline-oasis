module openline-ai/channels-api

go 1.19

require (
	github.com/DusanKasan/parsemail v1.2.0
	github.com/caarlos0/env/v6 v6.10.1
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.8.1
	github.com/gorilla/websocket v1.5.0
	github.com/openline-ai/openline-customer-os/packages/server/message-store v0.0.0-20221115083645-5dd530e6e4ab
	github.com/stretchr/testify v1.8.0
	github.com/xhit/go-simple-mail/v2 v2.12.0
	golang.org/x/net v0.1.0
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
	openline-ai/oasis-api v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-test/deep v1.0.8 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace openline-ai/oasis-api => ./../oasis-api
