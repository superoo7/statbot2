package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/joho/godotenv"
)

type list []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Whitelist get from whitelist.json
var Whitelist []string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	setupWhitelist()
}

func setupWhitelist() {
	file, err := ioutil.ReadFile("config/whitelist.json")
	if err != nil {
		log.Fatal("whitelist.json not exist")
	}
	var whitelists list
	_ = json.Unmarshal([]byte(file), &whitelists)
	if err != nil {
		log.Fatal("whitelist.json not exist")
	}
	for _, c := range whitelists {
		Whitelist = append(Whitelist, c.ID)
	}
}
