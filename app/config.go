package app

type ServerConfig struct {
	Port         int `envconfig:"SERVER_PORT" default:"8080"`
	ReadTimeout  int `envconfig:"SERVER_READ_TIMEOUT" default:"5"`
	WriteTimeout int `envconfig:"SERVER_WRITE_TIMEOUT" default:"10"`
	IdleTimeout  int `envconfig:"SERVER_IDLE_TIMEOUT" default:"120"`
	JwtIssuer string `envconfig:"SERVER_JWT_ISSUER" required:"true"`
}
