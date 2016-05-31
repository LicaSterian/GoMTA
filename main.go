package main

import (
	"io/ioutil"
	"log"
	"./utils"
	"./mta"
)

func main() {
	file, err := ioutil.ReadFile("./email/body.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	users := []mta.Recipient{
		mta.Recipient{Name: "Lica", Email:"lica.sterian@yahoo.com"},
		//mta.Recipient{Name: "Alexandra", Email:"bebesha89@yahoo.com"},
	}
	var m = mta.Mta{"mta5.am0.yahoodns.net"}
	for _, user := range(users) {
		parsedMessage := utils.ParseMessage(file, user)
		m.Send(parsedMessage)
	}
}