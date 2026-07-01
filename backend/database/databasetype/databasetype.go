package databasetype

// DBType defines the type of databse connected
type DBType int

const (
	// Sqlite driver
	Sqlite DBType = iota //0
	// MySQL driver
	MySQL // 1
	// PostgreSQL driver
	PostgreSQL // 2
)
