package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/go-vgo/robotgo"
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
		_, msg, err := c.ReadMessage()
		command := strings.Split(string(msg), " ")
		if err != nil {
			log.Println("Read: ", err)
			break
		}
		log.Printf("%s", msg)
		switch command[0] {
		case "MOVE_MOUSE":
			x0, _ := strconv.Atoi(command[1])
			y0, _ := strconv.Atoi(command[2])
			x0 = x0/5
			y0 = y0/5
			x, y := robotgo.GetMousePos()
			robotgo.MoveMouse(x + x0, y + y0)
		case "LEFT_CLICK":
			robotgo.MouseClick()
		}
	}
}

func main() {
	log.Printf("Starting Server")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", recieveCommand)
	log.Fatal(http.ListenAndServe(*addr, nil))
}