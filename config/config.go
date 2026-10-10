package config

import (
	"errors"
	"fmt"
	"io/fs"
	"net"
	"net/url"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	ENV      string `env:"ENV" envDefault:"local"`
	HTTP     HTTP
	Database Database
}

type HTTP struct {
	Addr         string        `env:"HTTP_ADDR" envDefault:"8080"`
	ReadTimeout  time.Duration `env:"HTTP_READTIMEOUT" envDefault:"10s"`
	WriteTimeout time.Duration `env:"HTTP_WRITETIMEOUT" envDefault:"10s"`
}

type Database struct {
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Port     string `env:"DATABASE_PORT" envDefault:"5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
	SSLMode  string `env:"DATABASE_SSLMODE" envDefault:"disabled"`

	MaxOpenConnections string        `env:"DATABASE_MAX_OPEN_CONNECTIONS" envDefault:"15"`
	MaxIdleConnections string        `env:"DATABASE_MAX_IDLE_CONNECTIONS" envDefault:"5"`
	ConnectionTimeout  time.Duration `env:"DATABASE_CONNECTION_TIMEOUT" envDefault:"15m"`
}

func (d Database) url() string {

	url := url.URL{
		Scheme:   "postgres",
		Host:     net.JoinHostPort(d.Host, d.Port),
		User:     url.UserPassword(d.Username, d.Password),
		Path:     d.Name,
		RawQuery: "sslmode=" + d.SSLMode,
	}

	return url.String()
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return &Config{}, fmt.Errorf("Error loading env : %v", err)
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return &Config{}, fmt.Errorf("Error parsing env : %v", err)
	}

	return &cfg, nil
}
