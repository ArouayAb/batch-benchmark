package models

import "time"

type Operation_type string

const (
	IN  Operation_type = "IN"
	OUT Operation_type = "OUT"
)

type Transaction struct {
	Operation_code uint `gorm:"primary_key"`
	Client_id      uint
	Client         Client `gorm:"foreignKey:Client_id"`
	Amount         float64
	Operation_type Operation_type
	Operation_date time.Time
}
