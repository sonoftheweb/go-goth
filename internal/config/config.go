package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DatabaseHost     string `envconfig:"DATABASE_HOST" default:"localhost"`
	DatabasePort     string `envconfig:"DATABASE_PORT" default:"5432"`
	DatabaseUser     string `envconfig:"DATABASE_USER" default:"postgres"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	DatabaseName     string `envconfig:"DATABASE_NAME" default:"postgres"`
	DatabaseSSLMODE  string `envconfig:"DATABASE_SSLMODE" default:"disable"`

	SessionCookieName string `envconfig:"SESSION_COOKIE_NAME" default:"session"`
	Port              string `envconfig:"PORT" default:":4000"`
}

func loadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
