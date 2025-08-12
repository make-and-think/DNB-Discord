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

	RedisClientFileHash struct {
		Uri string

		ReadBufferSize  int64
		WriteBufferSize int64
	}

	TaggerBackend struct {
		Url       string
		Token     string
		ImageSize int64
		ImageType string //TODO change to lib type of what used
	}

	Another struct {
		Test string
	}
}

// Global is the global configuration instance.
var Global Config

func LoadConfig(path string) (*Config, error) {
	if _, err := toml.DecodeFile(path, &Global); err != nil {
		return nil, err
	}
	return &Global, nil
}
