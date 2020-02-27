package core

import (
	"database/sql"
	"fmt"
)

//TODO: to config
var DataSourceName string = "root:gotest@/cms"
var Driver string = "mysql"

func GetDB() *sql.DB {
	db, err := sql.Open(Driver, DataSourceName)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
