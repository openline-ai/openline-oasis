package config

type Config struct {
	Postgres struct {
		Host            string `env:"POSTGRES_HOST,required"`
		Port            string `env:"POSTGRES_PORT,required"`
		User            string `env:"POSTGRES_USER,required,unset"`
		Db              string `env:"POSTGRES_DB,required"`
		Password        string `env:"POSTGRES_PASSWORD,required,unset"`
		MaxConn         int    `env:"POSTGRES_DB_MAX_CONN"`
		MaxIdleConn     int    `env:"POSTGRES_DB_MAX_IDLE_CONN"`
		ConnMaxLifetime int    `env:"POSTGRES_DB_CONN_MAX_LIFETIME"`
	}
	Service struct {
		MessageStoreUrl    string `env:"MESSAGE_STORE_URL,required"`
		MessageStoreApiKey string `env:"MESSAGE_STORE_API_KEY,required"`
		ServerAddress      string `env:"CHANNELS_API_SERVER_ADDRESS,required"`
		GRPCPort           int    `env:"CHANNELS_GRPC_PORT,required"`
		OasisApiUrl        string `env:"OASIS_API_URL,required"`
		CorsUrl            string `env:"CHANNELS_API_CORS_URL,required"`
	}
	Mail struct {
		ApiKey string `env:"MAIL_API_KEY,required"`
	}
	GMail struct {
		ClientId          string `env:"GMAIL_CLIENT_ID,unset"`
		ClientSecret      string `env:"GMAIL_CLIENT_SECRET,unset"`
		RedirectUris      string `env:"GMAIL_REDIRECT_URIS"`
		JavascriptOrigins string `env:"GMAIL_JAVASCRIPT_ORIGINS"`
	}
	WebChat struct {
		ApiKey          string `env:"WEBCHAT_API_KEY,required"`
		SlackWebhookUrl string `env:"SLACK_WEBHOOK_URL"`
		PingInterval    int    `env:"WEBSOCKET_PING_INTERVAL"`
	}
}
