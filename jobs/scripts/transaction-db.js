var { createPool } = require('mysql2');
var pool = createPool({
  host     : process.env.DB_URL,
  user     : process.env.DB_USER,
  password : process.env.DB_PASSWORD,
  database : process.env.DB_NAME,
  waitForConnections: true,
  connectionLimit: 10,
  queueLimit: 0
});

(async () => {
  let transactions = await new Promise((resolve) => { 
      pool.query(
        `SELECT * from transactions WHERE DATE(operation_date) > '${((yesterdayDate) => {
          return `${new Date(yesterdayDate
            .setDate(yesterdayDate.getDate() - 1))
            .toISOString()
            .slice(0, 19)
            .replace('T', ' ')}';`
        })(new Date)}`, 
        function (error, results, fields) {
          if (error) throw error;
          console.log(results);
          resolve(results);
        }
      );
  });

  let number_processed = 0;
  let number_transactions = transactions.length;

  await new Promise((resolve) => {
    for(const transaction of transactions) {
      pool.query(`UPDATE clients SET balance = balance ${transaction.operation_type === 'IN'? '+': '-'} ${transaction.amount} WHERE code = ${transaction.client_id}`, 
      function(error, results, fields) {
        if(error) throw error
        number_processed++;
      });
    }
    resolve();
  })

  console.log("::", number_processed / number_transactions, "%");

})();