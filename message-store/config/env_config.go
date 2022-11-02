package config

type Config struct {
	DB struct {
		Host     string `env:"DB_HOST,required"`
		Port     int    `env:"DB_PORT,required"`
		User     string `env:"DB_USER,required,unset"`
		Name     string `env:"DB_NAME,required"`
		Password string `env:"DB_PASSWORD,required,unset"`
	}
}
