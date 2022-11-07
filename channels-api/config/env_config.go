package config

type Config struct {
	Service struct {
		MessageStore  string `env:"MESSAGE_STORE_URL,required"`
		ServerAddress string `env:"CHANNELS_API_SERVER_ADDRESS,required"`
		GRPCPort      int    `env:"GRPC_PORT,required"`
	}
	Mail struct {
		SMTPSeverAddress  string `env:"SMTP_SERVER_ADDRESS,required"`
		SMTPSeverUser     string `env:"SMTP_SERVER_USER,required"`
		SMTPSeverPassword string `env:"SMTP_SERVER_PASSWORD,required"`
	}
}
