package api

import (
	"context"

	"github.com/aholake/payment-service/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
