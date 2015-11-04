package mta

import (
	"log"
	"net"
	"encoding/hex"
	"io"
	"time"
)

type Mta struct{
	Host string// = "mta5.am0.yahoodns.net"
}

func (m Mta) Send(messageBody MessageBody, rcptTo []string) {
	for _, to := range rcptTo {
		sendMessage(m.Host, messageBody, to)
	}
}

func sendMessage(host string, messageBody MessageBody, to string) {
	log.Println("sendMessage", host, to)
	conn, err := net.Dial("tcp", host + ":smtp")
	if err != nil {
		log.Fatal("err ", err)
	}
	onResponse(conn)
	for i, command := range messageBody.Data {
		time.Sleep(1 * time.Second)
		// TODO change hardcoded values with messageBody.
		sendCommand(conn, command, i < 4 || i >= 9)
	}
}

func sendCommand(conn net.Conn, command []byte, hasResponse bool) {
	log.Println("sendCommand", string(command), hasResponse)
	conn.Write(command)
	if hasResponse {
		time.Sleep(1 * time.Second)
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
	}
	log.Println(hex.Dump(data[:n]))
}