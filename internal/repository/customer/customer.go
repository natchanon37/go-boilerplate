package repository_customer

import (
	"context"
	"go-boilerplate/internal/repository/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	WithTx(txHandler *gorm.DB) CustomerRepository
	Create(ctx context.Context, cus *models.Customer) error
	GetById(ctx context.Context, cusId string) (*models.Customer, error)
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

func (cusRepo *customerRepository) GetById(ctx context.Context, cusId string) (mc *models.Customer, err error) {
	if err := cusRepo.preload().WithContext(ctx).Where("customer_id = ?", cusId).First(&mc).Error; err != nil {
		return nil, err
	}
	return mc, nil
}

func (cusRepo *customerRepository) Create(ctx context.Context, cus *models.Customer) error {
	return cusRepo.conn.WithContext(ctx).Create(&cus).Error
}

func (cusRepo *customerRepository) preload() *gorm.DB {
	return cusRepo.conn.Preload("Transactions")
}

func NewCustomerRepository(conn *gorm.DB) CustomerRepository {
	return &customerRepository{
		conn: conn,
	}
}
