package sqlite

type Sqlite interface {
	InsertUser(email, password string) error
}
