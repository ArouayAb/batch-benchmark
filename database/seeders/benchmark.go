package seeders

import (
	"batch/main/models"
	"time"

	"gorm.io/gorm"
)

func Benchmark(dbCon *gorm.DB) {
	var transaction = models.Transaction{
		OperationCode: 1,
		ClientID:      1,
		Amount:        43.2,
		OperationType: "IN",
		OperationDate: time.Now(),
	}

	var client = models.Client{
		Code:        1,
		Name:        "test",
		Balance:     423.523,
		Mail:        "sdcoiahu@oihsdf.csd",
		Transaction: []models.Transaction{transaction},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dbCon.Create(&client)
}
