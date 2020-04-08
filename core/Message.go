package core

type Message struct {
	Status  int    `json:"status"`
	Message Result `json:"message"`
}
