package ws

import (
	"encoding/json"
	"log"
)

// HelloMessage A simple struct that contains the message
type HelloMessage struct {
	// Message parameter
	Message string `json:"message"`
}

const (
	// HelloW Default message
	HelloW = "Hello, World!"
)

var hm = HelloMessage{Message: HelloW}

func Hello() string {
	m, err := json.Marshal(hm)
	if err != nil {
		log.Panic("error while trying to marshal. %s", err.Error())
		return ""
	}
	log.Printf("> %s", string(m))
	return string(m)
}

func HelloWorld() string {
	log.Printf("> HelloWorld")
	return "Hello, World!"
}
