package api

import (
	"context"

	"github.com/aholake/payment-service/internal/application/core/domain"
	"github.com/aholake/payment-service/internal/ports/db"
)

type Application struct {
	dbPort db.DBPort
}

func NewApplication(db db.DBPort) *Application {
	return &Application{dbPort: db}
}

func (a Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	if err := a.dbPort.Save(ctx, &payment); err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}
