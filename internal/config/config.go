package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Discord struct {
		Token string
	}
	RedisClientGuild struct {
		Uri string

		ReadBufferSize  int64
		WriteBufferSize int64
	}

	BackendEndpoints struct {
		Urls  []string
		Token string
	}

	Another struct {
		Test string
	}
}

// Values is the global configuration instance.
var Values Config

func LoadConfig(path string) (*Config, error) {
	if _, err := toml.DecodeFile(path, &Values); err != nil {
		return nil, err
	}
	return &Values, nil
}
