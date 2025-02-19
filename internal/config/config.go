package config

type (
	EnvVars struct {
		ServerHost string `envconfig:"server_host"`
		ServerPort int    `envconfig:"server_port"`
	}

	serverConfig struct {
		Host string
		Port int
	}

	Config struct {
		Server serverConfig
	}
)

func NewConfig(f EnvVars) Config {
	return Config{
		Server: serverConfig{
			Host: f.ServerHost,
			Port: f.ServerPort,
		},
	}
}
