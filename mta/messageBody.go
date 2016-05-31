package mta

type MessageBody struct {
	DataCommandIndex int
	EndDataCommandIndex int
	Data [][]byte
}