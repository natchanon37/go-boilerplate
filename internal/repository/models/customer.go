package models

import (
	"go-boilerplate/internal/repository/common"
)

type Customer struct {
	CustomerId   string        `gorm:"type:varchar(64);primaryKey" json:"customer_id"`
	Name         string        `gorm:"type:varchar(255)" json:"name"`
	Transactions []Transaction `gorm:"foreignKey:CustomerId" json:"transactions"`
	common.Timestamp
}
