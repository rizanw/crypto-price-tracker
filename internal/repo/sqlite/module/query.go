package module

const (
	qInsertUser = `INSERT INTO users VALUES(NULL,?,?) RETURNING id;`

	qFindUser = `SELECT * FROM users WHERE email=?;`

	qGetCoins = `SELECT * FROM coins WHERE user_id=?;`

	qInsertCoin = `INSERT INTO coins VALUES(NULL,?,?);`

	qDeleteCoin = `DELETE FROM coins WHERE user_id=? AND coin_id=?;`
)
