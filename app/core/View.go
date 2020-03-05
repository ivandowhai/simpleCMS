package core

import (
	"fmt"
	"html/template"
)

func GetView(name string) *template.Template {
	// TODO: move path to config
	templ, err := template.ParseFiles("templates/" + GetSettings().Template + "/" + name + ".html")
	if err != nil {
		fmt.Println(err)
	}

	return templ
}
