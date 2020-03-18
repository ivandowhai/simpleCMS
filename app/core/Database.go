package core

import (
	"database/sql"
)

func GetDB() *sql.DB {
	logger := Logger{}
	logger.Init()
	var settings = GetSettings()
	var dataSourceName = settings.DBUser + ":" + settings.DBPassword + "@" + settings.DBHost + "/" + settings.DBName

	db, err := sql.Open(settings.DBDriver, dataSourceName)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	return db
}
