package config

import (
	"log"
	"os"

	"crucial/banner/libs/util"

	"crucial/banner/app/flag"
)

const configPath = "./configs/"

var cfg Config

// Peek - возврашает инстанс конфига
func Peek() Config {
	return cfg
}

func init() {
	var (
		cfgFile = configPath + flag.Peek().ConfigFile
		err     error
	)

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(file)

	err = util.ReadJSON(cfgFile, &cfg)
	if err != nil {
		log.Panicln("config: " + err.Error())
	}

	log.Printf("config: parsed from %+s\n", cfgFile)
}

// Config - структура конфигураций
type Config struct {
	App      app      `json:"app"`
	Server   server   `json:"server"`
	Database database `json:"database"`
}

type app struct {
	Name   string `json:"name"`
	Stage  string `json:"stage"`
	APIKey string `json:"apiKey"`
}

type server struct {
	Port   string `json:"port"`
	Prefix string `json:"prefix"`
	URL    string `json:"url"`
}

type database struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"pass"`
	DBName   string `json:"dbName"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}
