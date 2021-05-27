package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ervices/ws"
)

const (
	//DefaultPort is the port on which the server will listen too
	DefaultPort = 8181
)

var port int

func init() {
	port = DefaultPort
	if p := strings.TrimSpace(os.Getenv("WS_GO_PORT")); p != "" {
		port, err := strconv.Atoi(p)
		if err == nil {
			log.Println(" port WS_GO_PORT = ", port)
		} else {
			log.Println(" defaulting to the port 8181")
		}
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/hello invoked")
	fmt.Fprintf(w, "%s", ws.Hello())
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/time invoked")
	fmt.Fprintf(w, "%s", ws.Time())
}

func main() {
	ws.Hello()
	fmt.Println(ws.HelloWorld())

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/time", timeHandler)
	log.Println(fmt.Sprintf(":%d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
