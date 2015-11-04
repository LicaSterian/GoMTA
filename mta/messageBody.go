package mta

type MessageBody struct {
//	Dns string
//	MailFrom string
	DataCommandIndex int
	EndDataCommandIndex int
	Data [][]byte
}