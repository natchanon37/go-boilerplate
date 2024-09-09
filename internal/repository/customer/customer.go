package repository_customer

import (
	"context"
	"go-boilerplate/internal/repository/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	WithTx(txHandler *gorm.DB) CustomerRepository
	Create(ctx context.Context, cus *models.Customer) error
}

type customerRepository struct {
	conn *gorm.DB
}

func (cusRepo *customerRepository) WithTx(txHandler *gorm.DB) CustomerRepository {
	if txHandler == nil {
		return cusRepo
	}
	cusRepo.conn = txHandler
	return cusRepo
}

func (cusRepo *customerRepository) Create(ctx context.Context, cus *models.Customer) error {
	return cusRepo.conn.WithContext(ctx).Create(&cus).Error
}

func NewCustomerRepository(conn *gorm.DB) CustomerRepository {
	return &customerRepository{
		conn: conn,
	}
}
