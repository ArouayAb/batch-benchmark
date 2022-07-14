package models

import "time"

type Client struct {
	Code        uint `gorm:"primaryKey"`
	Name        string
	Balance     float64
	Mail        string
	Transaction []Transaction `gorm:"foreignKey:ClientID;references:Code"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
