package utils

import (
	"./../mta"
	"log"
	"bufio"
	"bytes"
)

const newLine string = "\r\n"

func ParseMessage(file []byte, users []string) mta.MessageBody {
	reader := bytes.NewReader(file)
	scanner := bufio.NewScanner(reader)
	var data [][]byte
	var i = 0
	for scanner.Scan() {
		switch i {
			case 2:
				for _, user := range(users) {
					data = append(data, []byte("RCPT TO: <" + user + ">" + newLine))
				}
				data = append(data, []byte(scanner.Text() + newLine))
				break
			default:
				data = append(data, []byte(scanner.Text() + newLine))
				break
		}
//		log.Println(string(data[i]))
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	messageBody := mta.MessageBody{
		Data: data,
	}
	return messageBody
}