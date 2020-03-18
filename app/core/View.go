package core

import (
	"html/template"
)

func GetView(name string, layout string) *template.Template {
	logger := Logger{}
	logger.Init()
	templ, err := template.New("").
		ParseFiles("templates/"+GetSettings().Template+"/"+name+".html", "templates/"+GetSettings().Template+"/layout/"+layout+".html")
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	return templ
}
