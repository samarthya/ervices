package ws

import (
	"encoding/json"
	"log"
	"time"
)

// HelloMessage A simple struct that contains the message
type HelloMessage struct {
	// Message parameter
	Message string `json:"message"`
}

//TimeMessag A simple struct for the time message
type TimeMessage struct {
	// Message parameter
	MyTime time.Time `json:"time"`
}

const (
	// HelloW Default message
	HelloW = "Hello, World!"
)

var hm = HelloMessage{Message: HelloW}
var tm = TimeMessage{MyTime: time.Now()}

func Hello() string {
	m, err := json.Marshal(hm)
	if err != nil {
		log.Printf("error while trying to marshal. %s", err.Error())
		return ""
	}
	log.Printf("> %s", string(m))
	return string(m)
}

func Time() string {
	tm.MyTime = time.Now()
	m, err := json.Marshal(tm)
	if err != nil {
		log.Printf("error while trying to marshal. %s", err.Error())
		return ""
	}
	log.Printf("> %s", string(m))
	return string(m)
}

func HelloWorld() string {
	log.Printf("> HelloWorld")
	return "Hello, World!"
}
