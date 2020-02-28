package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Settings struct {
	DBDriver   string `json:"db_driver"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBName     string `json:"db_name"`
}

func GetSettings() *Settings {
	file, err := os.Open("config/main.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	settings := new(Settings)
	err = json.Unmarshal(b, &settings)
	if err != nil {
		log.Fatal(err)
	}

	return settings
}
