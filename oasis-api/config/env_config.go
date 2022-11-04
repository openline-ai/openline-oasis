package config

type Config struct {
	Service struct {
		MessageStore  string `env:"MESSAGE_STORE_URL,required"`
		CorsUrl       string `env:"CORS_URL,required"`
		ServerAddress string `env:"OASIS_API_SERVER_ADDRESS,required"`
	}
}
