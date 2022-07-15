package main

import (
	"batch-benchmark/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bitfield/script"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Transaction struct {
	Code   uint
	Amount float64
	OpType string
}

func main() {
	var transactions []Transaction

	transactionsFromJson, errIo := script.File("logfiles/transactions.json").String()
	if errIo != nil {
		log.Fatal(errIo)
	}

	json.Unmarshal([]byte(transactionsFromJson), &transactions)

	fmt.Println(transactions)

	var dsn = "root:password@tcp(localhost:3306)/batch_benchmark_db?charset=utf8mb4&parseTime=True&loc=Local"

	var db, errDb = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDb != nil {
		log.Fatal(errDb)
	}

	for _, transaction := range transactions {
		client := models.Client{}
		client.Code = transaction.Code
		db.First(&client)
		if transaction.OpType == "IN" {
			client.Balance += transaction.Amount
		}
		if transaction.OpType == "OUT" {
			client.Balance -= transaction.Amount
		}
		db.Save(&client)
	}
}
