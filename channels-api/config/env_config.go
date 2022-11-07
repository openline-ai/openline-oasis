package config

type Config struct {
	Service struct {
		MessageStore  string `env:"MESSAGE_STORE_URL,required"`
		ServerAddress string `env:"CHANNELS_API_SERVER_ADDRESS,required"`
	}
}
