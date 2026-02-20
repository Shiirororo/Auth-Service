package service

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
This this things before fixing anything in this file:

The token have following format:
*/
type RedisBlacklist struct {
	client *redis.Client
}

func NewRedisBlacklist(client *redis.Client) *RedisBlacklist {
	return &RedisBlacklist{client: client}
}

type TokenBlacklist interface {
	IsBlacklisted(ctx context.Context, jti string) (bool, error)
	SetRefreshToken(ctx context.Context, userID string, jti string, ttl time.Duration) error
}

func (r *RedisBlacklist) IsBlacklisted(ctx context.Context, jti string) (bool, error) {
	key := "TOKEN_BLACK_LIST" + jti
	exists, err := r.client.Exists(ctx, key).Result()
	return exists == 1, err
}
func (r *RedisBlacklist) SetRefreshToken(ctx context.Context, userID string, jti string, ttl time.Duration) error {
	key := "TOKEN_BLACK_LIST" + jti
	val := userID
	return r.client.Set(ctx, key, val, ttl).Err()
} // Write to jti
