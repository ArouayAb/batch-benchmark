package models

import (
	"time"
)

type Client struct {
	Code          uint `gorm:"primary_key"`
	Name          string
	Balance       float64
	Mail          string
	Creation_date time.Time
	Update_date   time.Time
}
