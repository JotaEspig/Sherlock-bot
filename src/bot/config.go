package bot

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	TOKEN  string
	PREFIX string
)

type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func init() {
	log.Println("Loading config file...")
	content, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Fatalln(err.Error())
	}

	var botConfig Config

	err = json.Unmarshal(content, &botConfig)

	if err != nil {
		log.Fatalln(err.Error())
	}

	TOKEN, PREFIX = botConfig.Token, botConfig.Prefix
	log.Print("Config file loaded!\n\n")
}
