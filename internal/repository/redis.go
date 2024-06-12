package repository

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/shankssc/Stripe-microservice/internal/model"
)

type RedisRepo struct {
	Client *redis.Client
}

func (r *RedisRepo) Insert(ctx context.Context, payment model.PaymentMethod) error {
	return nil
}
