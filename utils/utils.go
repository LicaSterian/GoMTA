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
	var dataCommandIndex int
	var endDataCommandIndex int
	var data [][]byte
	var i = 0
	for scanner.Scan() {
		text := scanner.Text() + newLine
		switch i {
			case 2:
				for _, user := range(users) {
					data = append(data, []byte("RCPT TO: <" + user + ">" + newLine))
				}
				data = append(data, []byte(text))
				dataCommandIndex = i + len(users)
				break
			default:
				if scanner.Text() == "." {
					text = newLine + "." + newLine
				}
				data = append(data, []byte(text))
				endDataCommandIndex = i + len(users) - 1
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