package models

import "time"

type Transaction struct {
	OperationCode uint `gorm:"primaryKey"`
	ClientID      uint
	Amount        float64
	OperationType string
	OperationDate time.Time
}
