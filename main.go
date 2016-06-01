package main

import (
	"./utils"
	"./mta"
)

func main() {
	path := "./email/mail.tpl"

	sender := mta.Sender{
		Domain: "mail.dom.eu",
		MailFrom: "<dirdel@tatutzu.com>",
		Header: mta.MailHeader{
			FriendlyFrom: `"dirdel" <dirdel@tatutzu.com>`,
			Subject: "Hello Lica. ",
		},
	}

	users := []mta.Recipient{
		mta.Recipient{Name: "Lica", Email:"lica.sterian@yahoo.com"},
		//mta.Recipient{Name: "Alexandra", Email:"bebesha89@yahoo.com"},
	}
	var m = mta.Mta{"mta5.am0.yahoodns.net"}
	for _, user := range(users) {
		parsedMessage := utils.ParseTemplate(path, &sender, user)
		m.Send(parsedMessage)
	}
}