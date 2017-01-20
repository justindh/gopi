package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

// Config stucture for the EVEData App
type Config struct {
	UserAgent string
	Database  struct {
		Import string
		Driver string
		Spec   string
	}
	Store struct {
		Key    string
		Domain string
	}
	CREST struct {
		SSO struct {
			ClientID    string
			SecretKey   string
			RedirectURL string
		}
		Token struct {
			ClientID    string
			SecretKey   string
			RedirectURL string
		}
		ESIAccessToken struct {
			ClientID    string
			SecretKey   string
			RedirectURL string

			AccessToken  string
			RefreshToken string
			Expiry       time.Time
			TokenType    string
		}
	}

	Redis struct {
		Address  string
		Password string
	}

	WebServer struct {
		Port string
	}
}

// ReadConfig should be run at startup and output shared between microservices via context.
func ReadConfig() (*Config, error) {
	c := Config{}

	if _, err := toml.DecodeFile("config/config.conf", &c); err != nil {
		return nil, err
	}
	return &c, nil
}
