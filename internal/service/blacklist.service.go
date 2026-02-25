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
	BlacklistSession(ctx context.Context, sessionID string, ttl time.Duration) error
	IsSessionBlacklisted(ctx context.Context, sessionID string) (bool, error)
	BlacklistJTI(ctx context.Context, jti string, ttl time.Duration) error
	IsJTIBlacklisted(ctx context.Context, jti string) (bool, error)
}

func (r *RedisBlacklist) BlacklistSession(ctx context.Context, sessionID string, ttl time.Duration) error {
	if ttl <= 0 {
		return nil
	}
	key := fmt.Sprintf("token:blacklist:session:%s", sessionID)
	return r.client.Set(ctx, key, "1", ttl).Err()
}

func (r *RedisBlacklist) IsSessionBlacklisted(ctx context.Context, sessionID string) (bool, error) {
	key := fmt.Sprintf("token:blacklist:session:%s", sessionID)
	exists, err := r.client.Exists(ctx, key).Result()
	return exists == 1, err
}

func (r *RedisBlacklist) BlacklistJTI(ctx context.Context, jti string, ttl time.Duration) error {
	if ttl <= 0 {
		return nil
	}
	key := fmt.Sprintf("token:blacklist:jti:%s", jti)
	return r.client.Set(ctx, key, "1", ttl).Err()
}

func (r *RedisBlacklist) IsJTIBlacklisted(ctx context.Context, jti string) (bool, error) {
	key := fmt.Sprintf("token:blacklist:jti:%s", jti)
	exists, err := r.client.Exists(ctx, key).Result()
	return exists == 1, err
}
