package core

import (
	"database/sql"
	"fmt"
)

func GetDB() *sql.DB {
	var settings *Settings = GetSettings()
	var dataSourceName string = settings.DB.User + ":" + settings.DB.Password + "@" + settings.DB.Host + "/" + settings.DB.Name

	db, err := sql.Open(settings.DB.Driver, dataSourceName)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
