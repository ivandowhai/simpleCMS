package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Settings struct {
	DB struct {
		Driver   string `json:"db.driver"`
		User     string `json:"db.user"`
		Password string `json:"db.password"`
		Host     string `json:"db.host"`
		Name     string `json:"db.name"`
	}
}

func GetSettings() *Settings {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	settings := new(Settings)
	err = json.Unmarshal(b, settings)
	if err != nil {
		log.Fatal(err)
	}

	return settings
}
