package core

import (
	"database/sql"
	"fmt"
)

func GetDB() *sql.DB {
	var settings = GetSettings()
	var dataSourceName = settings.DBUser + ":" + settings.DBPassword + "@" + settings.DBHost + "/" + settings.DBName

	db, err := sql.Open(settings.DBDriver, dataSourceName)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
