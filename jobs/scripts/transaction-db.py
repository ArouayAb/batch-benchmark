import os
import mysql.connector

print(os.getenv('DB_URL'),
  os.getenv('DB_USER'),
  os.getenv('DB_PASSWORD'),
  os.getenv('DB_NAME'))

db = mysql.connector.connect(
  host=os.getenv('DB_URL'),
  user=os.getenv('DB_USER'),
  password=os.getenv('DB_PASSWORD'),
  database=os.getenv('DB_NAME')
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