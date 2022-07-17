package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bitfield/script"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Transaction struct {
	Code   uint
	Amount float64
	OpType string
}

func main() {
	godotenv.Load(".env")
	var transactions []Transaction

	transactionsFromJson, errIo := script.File("jobs/batches/transactions.json").String()
	if errIo != nil {
		log.Fatal(errIo)
	}

	json.Unmarshal([]byte(transactionsFromJson), &transactions)

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var op string

	for _, transaction := range transactions {
		if transaction.OpType == "IN" {
			op = "+"
		} else {
			op = "-"
		}
		updateQuery := fmt.Sprintf("UPDATE clients SET balance = balance %s %f WHERE code = %d", op, transaction.Amount, transaction.Code)

		db.Query(updateQuery)
	}
}
