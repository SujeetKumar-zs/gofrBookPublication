package driver

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connection : makes the connection to the database
func Connection() (*sql.DB, error) {
	Db, err := sql.Open("mysql", "root:Sujeet@2001@tcp(127.0.0.1:3306)/library")
	if err != nil {
		return nil, err
	}

	return Db, nil
}
