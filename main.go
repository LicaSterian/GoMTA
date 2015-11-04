package main

import (
	"io/ioutil"
	"log"
	"./utils"
	"./mta"
)

func main() {
	file, err := ioutil.ReadFile("./body/body.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	users := []string{"lica.sterian@yahoo.com"}
	parsedMessage := utils.ParseMessage(file, users)
	var m = mta.Mta{"mta5.am0.yahoodns.net"}
	m.Send(parsedMessage, users)
}