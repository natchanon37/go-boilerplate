package models

import (
	"go-boilerplate/internal/repository/common"
)

type Transaction struct {
	TransactionID uint    `gorm:"primary_key" json:"transaction_id"`
	CustomerID    uint    `gorm:"type:uint" json:"customer_id"`
	Amount        float64 `gorm:"type:decimal(10,2)" json:"amount"`
	common.Timestamp
}
