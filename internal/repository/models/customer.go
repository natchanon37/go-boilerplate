package models

import (
	"go-boilerplate/internal/repository/common"
)

type Customer struct {
	ID           uint          `gorm:"primary_key" json:"id"`
	Name         string        `gorm:"type:varchar(255)" json:"name"`
	Transactions []Transaction `gorm:"foreignKey:CustomerID;references:ID" json:"transactions"`
	common.Timestamp
}
