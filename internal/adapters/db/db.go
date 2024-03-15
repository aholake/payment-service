package db

import (
	"context"
	"fmt"

	"github.com/aholake/payment-service/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderID    int64
	TotalPrice float32
}

type DBAdapter struct {
	db *gorm.DB
}

func NewDBAdapter(datasource string) (*DBAdapter, error) {
	db, err := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to DB, error: %v", err)
	}

	if err = db.AutoMigrate(&Payment{}); err != nil {
		return nil, fmt.Errorf("failed on auto migration, error: %v", err)
	}
	return &DBAdapter{
		db: db,
	}, nil
}

func (a DBAdapter) Get(ctx context.Context, id int64) (*domain.Payment, error) {
	var paymentEntity Payment
	a.db.First(&paymentEntity, id)
	return &domain.Payment{
		ID:         int64(paymentEntity.ID),
		CustomerID: paymentEntity.CustomerID,
		Status:     paymentEntity.Status,
		OrderID:    paymentEntity.OrderID,
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.Unix(),
	}, nil
}

func (a DBAdapter) Save(ctx context.Context, payment *domain.Payment) error {
	paymentEntity := Payment{
		OrderID:    payment.OrderID,
		CustomerID: payment.CustomerID,
		Status:     payment.Status,
		TotalPrice: payment.TotalPrice,
	}

	tx := a.db.Create(&paymentEntity)
	if tx.Error == nil {
		payment.ID = int64(paymentEntity.ID)
	}

	return tx.Error
}
