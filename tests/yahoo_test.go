package tests

import (
	"testing"
	"../mta"
	"../utils"
)

func TestYahoo(t *testing.T) {
	sender := mta.Sender{
		Domain: "mail.dom.eu",
		MailFrom: "<dirdel@tatutzu.com>",
		Header: mta.MailHeader{
			FriendlyFrom: `"dirdel" <dirdel@tatutzu.com>`,
			Subject: "Hello Lica. ",
		},
	}

	user := mta.Recipient{Name: "Lica", Email:"lica.sterian@yahoo.com"},
	var m = mta.Mta{"mta5.am0.yahoodns.net"}
	parsedMessage := utils.ParseTemplate("./email/mail.tpl", &sender, user)
	m.Send(parsedMessage)

	t.Log("check your yahoo mail account.")
}