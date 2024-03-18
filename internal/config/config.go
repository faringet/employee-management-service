package config

import (
	"fmt"

	env "github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
)

// Config contains ENV variables.
type Config struct {
	ServerHost     string `env:"SERVER_HOST"`
	ServerPort     int    `env:"SERVER_PORT" envDefault:"8080"`
	ServerLogLevel string `env:"SERVER_LOG_LEVEL" envDefault:"development"`

	PostgresConfig PostgresConfig `envPrefix:"POSTGRES_"`

	NatsConfig NatsConfig `envPrefix:"NATS_"`

	SMTPConfig SMTPConfig `envPrefix:"SMTP_"`

	S3Config S3Config `envPrefix:"S3_"`

	SwaggerConfig SwaggerConfig `envPrefix:"SWAGGER_"`

	OAuthConfig OAuthConfig `envPrefix:"OAUTH_"`

	TestEnv bool `env:"TEST_ENV" envDefault:"false"`

	Environment string `env:"ENVIRONMENT" envDefault:"development"`

	InviteURL string `env:"INVITE_URL"`
}

type NatsConfig struct {
	Host string `env:"HOST"`
	Port int    `env:"PORT"`
}

type SwaggerConfig struct {
	Host string `env:"HOST"`
}

type S3Config struct {
	EndPoint    string `env:"ENDPOINT" envDefault:"localhost:9000"`
	AccessKeyID string `env:"ACCESS_KEY_ID" envDefault:"admin"`
	SecretKey   string `env:"SECRET_ACCESS_KEY" envDefault:"admin1234"`
	BucketName  string `env:"BUCKET_NAME" envDefault:"test"`
	Secure      bool   `env:"SECURE" envDefault:"false"`
	Region      string `env:"REGION" envDefault:""`
}

func (n NatsConfig) ConnString() string {
	return fmt.Sprintf("ns://%s:%d", n.Host, n.Port)
}

// PostgresConfig contains ENV variables to configure PostgreSQL connection.
type PostgresConfig struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DB       string `env:"DB"`
}

func (p *PostgresConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", p.User, p.Password, p.Host, p.Port, p.DB)
}

type SMTPConfig struct {
	APIKey    string `env:"API_KEY"`
	FromEmail string `env:"FROM_EMAIL"`
}

type OAuthConfig struct {
	Connection   string   `env:"CONNECTION"`
	Domain       string   `env:"DOMAIN"`
	ClientID     string   `env:"CLIENT_ID"`
	ClientSecret string   `env:"CLIENT_SECRET"`
	Audiences    []string `env:"AUDIENCES" envSeparator:","`
}

// New settles ENV variables into Config structure.
func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("config: could not process config: %w", err)
	}

	return &cfg, nil
}
