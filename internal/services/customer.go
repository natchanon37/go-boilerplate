package service_customer

import (
	"context"
	"encoding/base64"
	"fmt"
	repository_customer "go-boilerplate/internal/repository/customer"
	"go-boilerplate/internal/repository/models"
	"go-boilerplate/pkg/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerService interface {
	WithTx(tx *gorm.DB) CustomerService
	CreateCustomer(ctx context.Context, req CreateCustomerReq) error
	GetCustomerById(ctx context.Context, cusId string) (*GetCustomerData, error)
}

type customerService struct {
	name    string
	cusRepo repository_customer.CustomerRepository
}

func (s customerService) WithTx(tx *gorm.DB) CustomerService {
	s.cusRepo = s.cusRepo.WithTx(tx)
	return s
}

func (s customerService) GetCustomerById(ctx context.Context, cusId string) (*GetCustomerData, error) {
	cusData, err := s.cusRepo.GetById(ctx, cusId)
	if err != nil {
		return nil, err
	}

	// decode amount
	decodedAmount, err := base64.StdEncoding.DecodeString(cusData.Transactions[0].Amount)
	if err != nil {
		return nil, err
	}
	realAmount, err := strconv.ParseFloat(string(decodedAmount), 64)
	if err != nil {
		return nil, err
	}

	// decode txId
	txIdStr := string(cusData.Transactions[0].TransactionId)
	oriTxId := txIdStr[:len(cusData.Transactions[0].TransactionId)-14]

	res := &GetCustomerData{
		TxId:   oriTxId,
		Amount: utils.ConvertTo2Decimal(realAmount),
	}

	return res, nil

}

func (s customerService) CreateCustomer(ctx context.Context, req CreateCustomerReq) error {
	amountStr := fmt.Sprintf("%v", req.Amount)
	enCusName := base64.StdEncoding.EncodeToString([]byte(req.CustomerName))
	endCusAmount := base64.StdEncoding.EncodeToString([]byte(amountStr))
	// gen uuid for customer id
	cusId := uuid.New().String()

	txId := fmt.Sprintf("%s%s", req.TxId, time.Now().Format("20060102150405"))

	txPayload := []models.Transaction{
		{
			CustomerId:    cusId,
			TransactionId: txId,
			Amount:        endCusAmount,
		},
	}

	cus := &models.Customer{
		CustomerId:   cusId,
		Name:         enCusName,
		Transactions: txPayload,
	}

	if err := s.cusRepo.Create(ctx, cus); err != nil {
		return err
	}

	return nil
}

func NewCustomerService(
	cusRepo repository_customer.CustomerRepository,
) CustomerService {
	name := "customerService"
	return &customerService{
		name:    name,
		cusRepo: cusRepo,
	}

}
