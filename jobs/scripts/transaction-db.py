import os
import mysql.connector
from datetime import datetime, timedelta

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
yesterday = datetime.today() - timedelta(days=1)
cursor.execute(f"SELECT * FROM transactions WHERE DATE(operation_date) > {yesterday}")
transactions = cursor.fetchall()

for transaction in transactions:
    if transaction[3] == 'IN':
        sql = "UPDATE clients SET balance = balance + %s WHERE code = %s"
    if transaction[3] == 'OUT':
        sql = "UPDATE clients SET balance = balance - %s WHERE code = %s"

    cursor.execute(sql, (transaction[2], transaction[1]))

    db.commit()