package ws

import "encoding/json"

// HelloMessage Aa simple structure that will have the hello message
type HelloMessage struct {
	//Message that needs to be returned
	Message string "json:message"
}

func HelloFromServer() ([]byte, error) {
	var helloMessage = HelloMessage{Message: "Hello from server!"}
	return json.Marshal(helloMessage)
}
