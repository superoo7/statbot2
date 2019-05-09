package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type list []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type versionFile struct {
	Version string `json:"version"`
}

// Version of statbot
var Version string

// Whitelist get from whitelist.json
var Whitelist []string

// SupportedCurrencies Supported Currency from CoinGecko
var SupportedCurrencies []string

// Env "development" or "production"
var Env string

// DiscordToken Token for discord bot
var DiscordToken string

func init() {
	fmt.Println("CONFIGURATING...")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	DiscordToken = os.Getenv("DISCORD_TOKEN")
	Env = os.Getenv("ENV")

	setupWhitelist()
	setupVersion()
	setupCurrency()
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

func setupVersion() {
	file, err := ioutil.ReadFile("config/version.json")
	if err != nil {
		log.Fatal("version.json not exist")
	}
	var v versionFile
	_ = json.Unmarshal([]byte(file), &v)
	if err != nil {
		log.Fatal("version.json not exist")
	}
	Version = v.Version
}

func setupCurrency() {
	file, err := ioutil.ReadFile("config/currency.json")
	if err != nil {
		log.Fatal("currency.json not exist")
	}
	var currencies map[string]bool
	_ = json.Unmarshal([]byte(file), &currencies)
	if err != nil {
		log.Fatal("currency.json not exist")
	}
	keys := reflect.ValueOf(currencies).MapKeys()
	for _, c := range keys {
		SupportedCurrencies = append(SupportedCurrencies, c.String())
	}
}
