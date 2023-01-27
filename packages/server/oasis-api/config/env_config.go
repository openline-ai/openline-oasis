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
		ChannelsApi        string `env:"CHANNELS_API_URL,required"`
		CorsUrl            string `env:"CORS_URL,required"`
		ServerAddress      string `env:"OASIS_API_SERVER_ADDRESS,required"`
		GRPCPort           int    `env:"OASIS_GRPC_PORT,required"`
	}
	WebRTC struct {
		AuthSecret   string `env:"WEBRTC_AUTH_SECRET,required"`
		TTL          int    `env:"WEBRTC_AUTH_TTL,required"`
		PingInterval int    `env:"WEBSOCKET_PING_INTERVAL"`
	}
}
