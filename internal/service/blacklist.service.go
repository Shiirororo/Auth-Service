package service

import (
	"context"
	"fmt"
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

func NewRedisBlacklist(client *redis.Client) TokenBlacklist {
	return &RedisBlacklist{client: client}
}

type TokenBlacklist interface {
	IsBlacklisted(ctx context.Context, userID string, jti string) (bool, error)
	RevokeRefreshToken(ctx context.Context, userID string, jti string, ttl time.Time) error
}

func (r *RedisBlacklist) IsBlacklisted(ctx context.Context, userID string, jti string) (bool, error) {
	key := "token:blacklist:" + userID + "_" + jti
	exists, err := r.client.Exists(ctx, key).Result()
	return exists == 1, err
}
func (r *RedisBlacklist) RevokeRefreshToken(ctx context.Context, userID string, jti string, exp time.Time) error {
	ttl := time.Until(exp)
	if ttl <= 0 {
		return nil
	}

	key := fmt.Sprintf("token:blacklist:%s", userID+"_"+jti)
	return r.client.Set(ctx, key, "1", ttl).Err()
} // Write to jti
