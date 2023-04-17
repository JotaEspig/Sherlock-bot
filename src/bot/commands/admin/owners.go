package admin

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type owner struct {
	Id string `json:"id"`
}

type owners struct {
	Users []owner `json:"owners"`
}

var Owners owners

func init() {
	content, err := ioutil.ReadFile("./src/bot/commands/admin/owners.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = json.Unmarshal(content, &Owners)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
