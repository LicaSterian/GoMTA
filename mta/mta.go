package mta

import (
	"log"
	"net"
	"encoding/hex"
	"io"
	"time"
)

var smtpHost string

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
		time.Sleep(2 * time.Second)
		sendCommand(conn, command, i < 3)
	}
}

func sendCommand(conn net.Conn, command []byte, hasResponse bool) {
	log.Println("sendCommand", string(command), hasResponse)
	conn.Write(command)
//	if !respondNow {

//	}
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
	}
	log.Printf("onResponse:\n", hex.Dump(data[:n]))
	log.Println("==============================")
}


/*
func sendMessage(host string, messageBody MessageBody, to string) {
	conn, err := net.Dial("tcp", host + ":smtp")
	if err != nil {
		log.Fatal("err ", err)
	}
	log.Printf("%+v", conn)
	log.Printf("%+v", messageBody)

*/

	/*if err := conn.Mail(messageBody.MailFrom); err != nil {
		log.Fatal(err)
	}
	if err := conn.Rcpt(to); err != nil {
		log.Fatal(err)
	}

	log.Println(host)
	log.Println(messageBody.MailFrom)
	log.Println(to)
	log.Println(messageBody.Body)

	// Send the email body.

	wc, err := conn.Data()
	if err != nil {
		log.Fatal(err)
	}

//	_, err = log.Printf(wc, getMessageBody(mailMessage))
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = conn.Quit()
	if err != nil {
		log.Fatal(err)
	}*/
//}

/*
func getMessageBody(mailMessage MailMessage) string {
	var newLine = "\r\n"
	var result string

	result += "from: " + mailMessage.from + newLine
	result += "to: " + mailMessage.to + newLine
	result += "subject: " + mailMessage.subject + newLine + newLine
	result += mailMessage.body + newLine

	return result
}
*/