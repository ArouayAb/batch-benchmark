import mysql.connector

db = mysql.connector.connect(
  host="localhost",
  user="root",
  password="password",
  database="batchbenchmark"
)

cursor = db.cursor()

cursor.execute("SELECT * FROM transactions")
transactions = cursor.fetchall()

for transaction in transactions:
    if transaction[3] == 'IN':
        sql = "UPDATE clients SET balance = balance + %s WHERE code = %s"
    if transaction[3] == 'OUT':
        sql = "UPDATE clients SET balance = balance - %s WHERE code = %s"

    cursor.execute(sql, (transaction[2], transaction[1]))

    db.commit()