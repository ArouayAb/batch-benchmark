package helpers

import (
	"batch-benchmark/models"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func SetupDB(db *gorm.DB) {
	db.AutoMigrate(&models.Client{}, &models.Transaction{})
	seed(db)
}

func seed(db *gorm.DB) {

	for i := 1; i <= 20; i++ {
		db.Create(&models.Client{
			Name:          "Client" + strconv.Itoa(i),
			Balance:       300 * float64(i),
			Mail:          "client" + strconv.Itoa(i) + "@mail.com",
			Creation_date: time.Now(),
			Update_date:   time.Now(),
		})
	}

	operationsTypes := []models.Operation_type{models.IN, models.OUT}

	for k := 1; k <= 20; k++ {
		for j := 1; j <= 3; j++ {
			db.Create(&models.Transaction{
				Client_id:      uint(k),
				Amount:         90 * float64(rand.Intn(3)+1),
				Operation_type: operationsTypes[rand.Intn(2)],
				Operation_date: time.Now(),
			})
		}
	}
}
