package module

const (
	qInsertUser = `INSERT INTO users VALUES(NULL,?,?);`

	qFindUser = `SELECT * FROM users WHERE email=?;`
)
