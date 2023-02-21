package config

type Config struct {
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
		ClientId     string `env:"GMAIL_CLIENT_ID,unset"`
		ClientSecret string `env:"GMAIL_CLIENT_SECRET,unset"`
		RedirectUris string `env:"GMAIL_REDIRECT_URIS"`
		OryApiKey    string `env:"ORY_API_KEY,unset"`
		OryServerUrl string `env:"ORY_SERVER_URL"`
	}
	WebChat struct {
		ApiKey          string `env:"WEBCHAT_API_KEY,required"`
		SlackWebhookUrl string `env:"SLACK_WEBHOOK_URL"`
		PingInterval    int    `env:"WEBSOCKET_PING_INTERVAL"`
	}
	VCon struct {
		ApiKey          string `env:"VCON_API_KEY,required"`
		AwsAccessKey    string `env:"AWS_ACCESS_KEY"`
		AwsAccessSecret string `env:"AWS_ACCESS_SECRET"`
		AwsRegion       string `env:"AWS_REGION"`
		AwsBucket       string `env:"AWS_BUCKET"`
	}
}
