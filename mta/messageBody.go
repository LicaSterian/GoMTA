package mta

import "github.com/gin-gonic/gin"

type MessageBody struct {
	DataCommandIndex int
	EndDataCommandIndex int
	Data [][]byte
}

type MailHeader struct {
	FriendlyFrom string
	Subject string
}

type Sender struct {
	Domain, MailFrom string
	Header MailHeader
	Body gin.H
}