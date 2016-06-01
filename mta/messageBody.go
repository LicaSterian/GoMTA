package mta

type MessageBody struct {
	DataCommandIndex int
	EndDataCommandIndex int
	Data [][]byte
}

type MailHeader struct {
	FriendlyFrom string
	Subject string
}

type Map map[string]interface{}

type Sender struct {
	Domain, MailFrom string
	Header MailHeader
	Body Map
}