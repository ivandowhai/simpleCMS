package core

import (
	"fmt"
	"html/template"
)

func GetView(name string, layout string) *template.Template {
	templ, err := template.New("").
		ParseFiles("templates/"+GetSettings().Template+"/"+name+".html", "templates/"+GetSettings().Template+"/layout/"+layout+".html")
	if err != nil {
		fmt.Println(err)
	}

	return templ
}
