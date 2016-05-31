package mta

import (
	"log"
	"net"
	"encoding/hex"
	"io"
//	"time"
)

type Mta struct{
	Host string
}

var currentMessageBody MessageBody
var currentHost string

func (m Mta) Send(messageBody MessageBody) {
	log.Println("sending message")
	currentHost = m.Host
	send(messageBody)
}

func send(messageBody MessageBody) {
	currentMessageBody = messageBody
	sendMessage(currentHost, messageBody)
}

func retrySend() {
	log.Println("retrying sending of message")
	send(currentMessageBody)
}

func sendMessage(host string, messageBody MessageBody) {
	log.Println("sendMessage", host)
	conn, err := net.Dial("tcp", host + ":smtp")
	if err != nil {
		log.Fatal("err ", err)
	}
	onResponse(conn)
	for i, command := range messageBody.Data {
		sendCommand(conn, command, i <= messageBody.DataCommandIndex || i >= messageBody.EndDataCommandIndex)
	}
}

func sendCommand(conn net.Conn, command []byte, hasResponse bool) {
	log.Println("sendCommand", string(command), hasResponse)
	conn.Write(command)
	if hasResponse {
		onResponse(conn)
	}
}

func onResponse(conn net.Conn) {
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		if err != io.EOF {
			panic(err)
		} else {
			log.Printf("err: %v", err)
		}
		retrySend()
	}
	log.Println(hex.Dump(data[:n]))
}