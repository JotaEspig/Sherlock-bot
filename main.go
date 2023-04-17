package main

import (
	"log"
	"sherlock-bot/src/bot"
)

func main() {

	err := bot.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}

	<-make(chan struct{})
}
