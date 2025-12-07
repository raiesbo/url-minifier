package db

import (
	"errors"

	"github.com/redis/go-redis/v9"
)

var (
	ErrMissingURI = errors.New("no connection URI provided")
)

func NewRedis(uri string) *redis.Client {
	if uri == "" {
		panic(ErrMissingURI)
	}

	options, err := redis.ParseURL(uri)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(options)
}
