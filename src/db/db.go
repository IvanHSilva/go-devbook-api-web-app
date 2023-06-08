package db

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

// Open connection with SQL Server Database
func DBConnect() (*sql.DB, error) {
	//connection := "'Server=SERVIDOR\\SQLSERVER;Database=eCommerce;Trusted Connection=true;'" //Initial Catalog=eCommerce;Integrated Security=True;Encrypt=False;"
	db, err := sql.Open(config.DBType, config.Connection)

	if err != nil {
		return nil, err
	}

	//fmt.Println(db)
	if err = db.Ping(); err != nil {
		db.Close()
	}

	fmt.Println("Conexão Ok!")
	fmt.Println(db.Stats().OpenConnections)
	return db, nil
}
