package core

import (
	"log"
	"os"
	"time"
)

type Logger struct {
	settings *Settings
}

func (l *Logger) Init() *Logger {
	l.settings = GetSettings()
	return l
}

func (l *Logger) WriteLog(message string, messageType string) {
	dir := l.settings.LogsDirectory
	date := time.Now()
	filename := dir + "/" + messageType + "" + date.Format("2006-02-01") + ".log"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	if _, err := file.Write([]byte(message + "\n")); err != nil {
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
