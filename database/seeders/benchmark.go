package seeders

import (
	"batch/main/models"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func Benchmark(dbCon *gorm.DB) {
	for i := 1; i <= 20; i++ {
		dbCon.Create(&models.Client{
			Name:      "Client" + strconv.Itoa(i),
			Balance:   500 * float64(i),
			Mail:      "client" + strconv.Itoa(i) + "@mail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	operationsTypes := []models.OperationType{models.IN, models.OUT}

	for k := 1; k <= 20; k++ {
		for j := 1; j <= 3; j++ {
			dbCon.Create(&models.Transaction{
				ClientID:      uint(k),
				Amount:        90 * float64(rand.Intn(3)+1),
				OperationType: operationsTypes[rand.Intn(2)],
				OperationDate: time.Now(),
			})
		}
	}
}
