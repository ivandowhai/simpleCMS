package core

import (
	"database/sql"
	"fmt"
)

func GetDB() *sql.DB {
	var settings *Settings = GetSettings()
	var dataSourceName string = settings.DBUser + ":" + settings.DBPassword + "@" + settings.DBHost + "/" + settings.DBName

	db, err := sql.Open(settings.DBDriver, dataSourceName)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
