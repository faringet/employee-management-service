package config

import (
	"fmt"

	env "github.com/caarlos0/env/v10"
)

// Config contains ENV variables.
type Config struct {
	ServerHost     string `env:"SERVER_HOST,notEmpty"`
	ServerPort     int    `env:"SERVER_PORT" envDefault:"8080"`
	ServerLogLevel string `env:"SERVER_LOG_LEVEL" envDefault:"development"`

	SwaggerURL string `env:"SWAGGER_URL,notEmpty"`

	PostgresConfig PostgresConfig `envPrefix:"POSTGRES_"`

	NatsConfig NatsConfig `envPrefix:"NATS_"`

	TestEnv bool `env:"TEST_ENV" envDefault:"false"`
}

// PgConfig contains ENV variables to configure PostgreSQL connection.
type PostgresConfig struct {
	Host     string `env:"HOST,notEmpty"`
	Port     int    `env:"PORT,notEmpty"`
	User     string `env:"USER,notEmpty"`
	Password string `env:"PASSWORD,notEmpty"`
	DB       string `env:"DB,notEmpty"`
}

func (p PostgresConfig) ConnString() string {
	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", p.User, p.DB, p.Password, p.Host, p.Port)
}

type NatsConfig struct {
	Host string `env:"HOST,notEmpty"`
	Port int    `env:"PORT,notEmpty"`
}

func (n NatsConfig) ConnString() string {
	return fmt.Sprintf("nats://%s:%d", n.Host, n.Port)
}

// NewConfig settles ENV variables into Config structure.
func NewConfig() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("config: could not process config: %w", err)
	}

	return &cfg, nil
}
