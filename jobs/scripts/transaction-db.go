package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Transaction struct {
	OperationCode uint
	ClientID      uint
	Amount        float64
	OperationType string
	OperationDate time.Time
}

func main() {
	fmt.Println("Hello")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/batchbenchmark?parseTime=true")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	yesterdayTime := time.Now().Add(-24 * time.Hour)
	sqlQuery := fmt.Sprintf("SELECT * from transactions WHERE DATE(operation_date) > \"%s\"", yesterdayTime)

	transactions, errQ := db.Query(sqlQuery)
	defer transactions.Close()

	if errQ != nil {
		log.Fatal(errQ)
	}

	var number_transactions int
	db.QueryRow("SELECT count(*) FROM transactions").Scan(&number_transactions)

	number_processed := 0
	for transactions.Next() {
		var transaction Transaction

		err := transactions.Scan(
			&transaction.OperationCode,
			&transaction.ClientID,
			&transaction.Amount,
			&transaction.OperationType,
			&transaction.OperationDate,
		)

		if err != nil {
			log.Fatal(err)
			break
		}

		var op string
		if transaction.OperationType == "IN" {
			op = "+"
		} else {
			op = "-"
		}

		updateQuery := fmt.Sprintf("UPDATE clients SET balance = balance %s %f WHERE code = %d", op, transaction.Amount, transaction.ClientID)

		db.Query(updateQuery)
		number_processed++
	}

	log.Println("::", (number_processed/number_transactions)*100, "%", "Completed")
}
