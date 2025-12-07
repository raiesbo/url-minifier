package url

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	db *redis.Client
}

func NewRedisRepository(c *redis.Client) *RedisRepository {
	return &RedisRepository{db: c}
}

// Save stores the URL
func (r *RedisRepository) Save(url *URL) error {
	bytes, err := json.Marshal(url)
	if err != nil {
		return err
	}
	ctx := context.Background()
	pipe := r.db.Pipeline()
	pipe.Set(ctx, url.Slug, bytes, 0)

	// Secondary index system
	originalSlug := fmt.Sprintf("idx:original:%s", url.Original)
	pipe.SAdd(ctx, originalSlug, url.Slug)

	_, err = pipe.Exec(ctx)
	return err
}

// FindByKey searches by URL key and returns the result
func (r *RedisRepository) FindByKey(key string) (*URL, error) {
	url := &URL{}
	val, err := r.db.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return url, json.Unmarshal([]byte(val), url)
}
