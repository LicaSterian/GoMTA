package utils

import (
	"./../mta"
	"log"
	"bufio"
	"bytes"
)

const newLine string = "\r\n"

// make commands from DATA to . all in one command and store it in a body
func ParseMessage(file []byte, recipient mta.Recipient) mta.MessageBody {
	reader := bytes.NewReader(file)
	scanner := bufio.NewScanner(reader)
	var dataCommandIndex int
	var endDataCommandIndex int
	var data [][]byte
	var i = 0
	for scanner.Scan() {
		text := scanner.Text() + newLine
		switch i {
			case 2:
				data = append(data, []byte("RCPT TO: <" + recipient.Email + ">" + newLine))
				i++
				data = append(data, []byte(text))
				dataCommandIndex = i
				data = append(data, []byte("To: " + recipient.Name + " <" + recipient.Email + ">" + newLine))
				i++
				break
			default:
				if scanner.Text() == "." {
					text = newLine + "." + newLine
					endDataCommandIndex = i
				}
				data = append(data, []byte(text))
				break
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	messageBody := mta.MessageBody{
		DataCommandIndex: dataCommandIndex,
		EndDataCommandIndex: endDataCommandIndex,
		Data: data,
	}
	return messageBody
}