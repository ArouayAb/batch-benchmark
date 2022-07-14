package database

import (
	"batch/main/database/seeders"
	"batch/main/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbCon *gorm.DB

func Init() {
	dsn := "root:password@tcp(127.0.0.1:3306)/BatchBenchmark?charset=utf8mb4&parseTime=True&loc=Local"
	DbCon, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DbCon.Migrator().DropTable(&models.Client{}, &models.Transaction{})
	DbCon.AutoMigrate(&models.Client{}, &models.Transaction{})

	seeders.Benchmark(DbCon)
}
