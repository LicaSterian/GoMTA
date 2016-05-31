package main

import (
	"io/ioutil"
	"text/template"
	"fmt"
	"github.com/gin-gonic/gin"
	"bytes"
	"./utils"
	"./mta"
)

func main() {
	file, err := ioutil.ReadFile("./email/mail.tpl")
	if err != nil {
		fmt.Println("ioutil.ReadFile error :", err.Error())
	}

	sender := mta.Sender{
		Domain: "mail.dom.eu",
		MailFrom: "<dirdel@tatutzu.com>",
		Header: mta.MailHeader{
			FriendlyFrom: `"dirdel" <dirdel@tatutzu.com>`,
			Subject: "Hello Lica. ",
		},
	}

	tpl, err := template.New("mailTemplate").Parse(string(file))
	if err != nil {
		fmt.Println("template.New error :", err.Error())
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, gin.H{
		"Sender": sender,
	})
	if err != nil {
		fmt.Println("template.Execute error :", err.Error())
	}

	users := []mta.Recipient{
		mta.Recipient{Name: "Lica", Email:"lica.sterian@yahoo.com"},
		//mta.Recipient{Name: "Alexandra", Email:"bebesha89@yahoo.com"},
	}
	var m = mta.Mta{"mta5.am0.yahoodns.net"}
	for _, user := range(users) {
		parsedMessage := utils.ParseMessage(buf.Bytes(), user)
		m.Send(parsedMessage)
	}
}