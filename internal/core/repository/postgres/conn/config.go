package core_repository_postgres_pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port     string        `envconfig:"PORT" default:"5432"`
	Host     string        `envconfig:"HOST" required:"true"`
	User     string        `envconfig:"USER" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Database string        `envconfig:"DB" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("POSTGRES", &config); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("create config: %w", err)
		panic(err)
	}
	return config

}
