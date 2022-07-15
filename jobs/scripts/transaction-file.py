import mysql.connector
import json

mydb = mysql.connector.connect(
  host="localhost",
  user="root",
  password="password",
  database="batchbenchmark"
)

transactions_file = open('jobs/batches/transactions.json')

transactions = json.load(transactions_file)

mycursor = mydb.cursor()
number_transactions = len(transactions)
number_processed = 0
try:
  for transaction in transactions:
      sql = f"UPDATE clients SET balance = balance {'+' if transaction['OpType']=='IN' else '-'} {transaction['Amount']} WHERE code = {transaction['Code']}"

      mycursor.execute(sql)

      mydb.commit()

      number_processed += 1
except Exception as Error:
  print(f"{Error=}")

print('::', (number_processed / number_transactions) * 100, '%', 'Completed')
  

