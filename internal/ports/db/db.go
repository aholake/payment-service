package db

import (
	"context"

	"github.com/aholake/payment-service/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id int64) (*domain.Payment, error)
	Save(ctx context.Context, payment *domain.Payment) error
}
