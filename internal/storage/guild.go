package storage

import (
	"DNB-Discord/internal/config"
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	guildRDB *redis.Client
	guildCtx = context.Background()
)

type Tag struct {
	name      string
	threshold float64
}

type GuildConfig struct {
	// The ID of the guild in which the message was sent.
	GuildID string

	Tags []map[string]float32

	Rating struct {
		general      float32 //safe
		sensitive    float32
		questionable float32
		explicit     float32
	}
}

func GuildConfigInit() {
	opts, err := redis.ParseURL(config.Values.RedisClientGuild.Uri)
	if err != nil {
		panic(err)
	}

	guildRDB = redis.NewClient(opts)

	// Ping the Redis server to check the connection.
	_, err = guildRDB.Ping(guildCtx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func getGuildID(guildID string) (*GuildConfig, error) {
	// Get the encoded data from Redis.
	val, err := guildRDB.Get(guildCtx, guildID).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("no config found for guild ID: %s", guildID)
		}
		return nil, fmt.Errorf("failed to get GuildConfig from Redis: %w", err)
	}
	var guildConfig GuildConfig

	network := bytes.NewBuffer(val)
	dec := gob.NewDecoder(network)
	if err := dec.Decode(&guildConfig); err != nil {
		return nil, fmt.Errorf("failed to decode GuildConfig: %w", err)
	}

	return &guildConfig, nil
}

func setGuildConfig(guildConfig GuildConfig) error {
	var network bytes.Buffer
	// Create a new gob encoder and encode the config struct.
	enc := gob.NewEncoder(&network)

	if err := enc.Encode(guildConfig); err != nil {
		return fmt.Errorf("failed to encode GuildConfig: %w", err)
	}

	// Set the encoded data in Redis with the GuildID as the key.
	err := guildRDB.Set(guildCtx, guildConfig.GuildID, network.Bytes(), 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set GuildConfig in Redis: %w", err)
	}
	return nil
}
