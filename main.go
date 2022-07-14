package main

import (
	"batch-benchmark/helpers"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var dsn = "root:password@tcp(localhost:3306)/batch_benchmark_db?charset=utf8mb4"

	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	helpers.SetupDB(db)
}
