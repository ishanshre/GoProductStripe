package driver

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenDB(driverName, dsn string) (*sql.DB, error) {
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}
