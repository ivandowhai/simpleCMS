package core

import "io/ioutil"

type Logger struct {
	settings *Settings
}

func (l *Logger) Init() *Logger {
	l.settings = GetSettings()
	return l
}

func (l *Logger) WriteLog(error string) {
	dir := l.settings.LogsDirectory
	ioutil.WriteFile(dir+"/error.log", []byte(error+"\n"), 0644)
}
