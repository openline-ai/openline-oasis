module github.com/openline-ai/openline-oasis/packages/server/channels-api

go 1.19

replace github.com/openline-ai/openline-oasis/packages/server/oasis-api => ./../oasis-api

//replace github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module => ../../../../openline-customer-os/packages/server/customer-os-common-module
//replace github.com/openline-ai/openline-customer-os/packages/server/message-store => ../../../../openline-customer-os/packages/server/message-store

require (
	github.com/caarlos0/env/v6 v6.10.1
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.8.2
	github.com/gorilla/websocket v1.5.0
	github.com/joho/godotenv v1.4.0
	github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module v0.0.0-20230127043444-f21dd24e0515
	github.com/openline-ai/openline-customer-os/packages/server/message-store v0.0.0-20230113175908-2175601d4942
	github.com/openline-ai/openline-oasis/packages/server/oasis-api v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.1
	github.com/xhit/go-simple-mail/v2 v2.13.0
	golang.org/x/net v0.4.0
	google.golang.org/grpc v1.52.3
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/go-test/deep v1.0.8 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.24.3 // indirect
)
