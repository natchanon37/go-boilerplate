package models

import (
	"go-boilerplate/internal/repository/common"
)

type Transaction struct {
	CustomerId    string `gorm:"primaryKey" json:"-"`
	TransactionId string `gorm:"type:varchar(64);primaryKey" json:"transaction_id"`
	Amount        string `gorm:"type:varchar(64)" json:"amount"`
	common.Timestamp
}
