package config

type Config struct {
	Service struct {
		MessageStore  string `env:"MESSAGE_STORE_URL,required"`
		ChannelsApi   string `env:"CHANNELS_API_URL,required"`
		CorsUrl       string `env:"CORS_URL,required"`
		ServerAddress string `env:"OASIS_API_SERVER_ADDRESS,required"`
	}
	WebRTC struct {
		AuthSecret string `env:"WEBRTC_AUTH_SECRET,required"`
		TTL        int    `env:"WEBRTC_AUTH_TTL,required"`
	}
}
