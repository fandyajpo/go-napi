package connection

type Database struct{}

type DatabaseInterface interface {
	ConnectDatabase()
	QueryExecutor()
}

func NewDB() *Database {
	return &Database{}
}

func (d *Database) ConnectDatabase() {}

func (d *Database) QueryExecutor() {}
