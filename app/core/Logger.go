package core

import (
	"io/ioutil"
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
	err := ioutil.WriteFile(dir+"/"+messageType+""+date.String()+".log", []byte(message+"\n"), 0644)
	if err != nil {
		panic(err)
	}
}
