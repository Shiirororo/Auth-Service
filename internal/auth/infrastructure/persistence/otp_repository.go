package persistence

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/user_service/internal/auth/domain/repository"
)

type redisOTPRepository struct {
	client *redis.Client
}

func NewRedisOTPRepository(client *redis.Client) repository.OTPRepository {
	return &redisOTPRepository{
		client: client,
	}
}

func (r *redisOTPRepository) getRedisKey(email string) string {
	return fmt.Sprintf("otp:%s", email)
}

func (r *redisOTPRepository) SaveOTP(ctx context.Context, email string, otp int, ttl time.Duration) error {
	key := r.getRedisKey(email)
	return r.client.Set(ctx, key, otp, ttl).Err()
}

func (r *redisOTPRepository) GetOTP(ctx context.Context, email string) (int, error) {
	key := r.getRedisKey(email)
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, fmt.Errorf("otp not found or expired")
		}
		return 0, err
	}

	otp, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return otp, nil
}

func (r *redisOTPRepository) DeleteOTP(ctx context.Context, email string) error {
	key := r.getRedisKey(email)
	return r.client.Del(ctx, key).Err()
}
