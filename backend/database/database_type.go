package database

// DatabaseType defines the type of databse connected
type DatabaseType int

const (
	Sqlite     DatabaseType = iota //0
	MySQL                          // 1
	PostgreSQL                     // 2
)
