package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "192.168.0.28:8080", "http service address")

var upgrader = websocket.Upgrader{}

func recieveCommand(w http.ResponseWriter, r *http.Request) {
	log.Printf("Connection")

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade: ", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Read: ", err)
			break
		}
		log.Printf("Message: %s", message)
		log.Printf("mt: %s", mt)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", recieveCommand)
	log.Fatal(http.ListenAndServe(*addr, nil))
}