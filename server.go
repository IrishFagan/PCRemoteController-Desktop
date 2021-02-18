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

type Coordinates struct {
	x int
	y int
}

func recieveCommand(w http.ResponseWriter, r *http.Request) {
	log.Printf("Connection")

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade: ", err)
		return
	}
	defer c.Close()
	for {
		mt, msg, err := c.ReadMessage()
		message := string(msg)
		coords := strings.Split(message, " ")
		if err != nil {
			log.Println("Read: ", err)
			break
		}
		log.Printf("%s", message)
		log.Printf(coords[0])
		log.Printf(coords[1])
		cx, err := strconv.Atoi(coords[0])
		cy, err := strconv.Atoi(coords[1])
		log.Printf("mt: %s", mt)
		x, y := robotgo.GetMousePos()
		robotgo.MoveMouse(x + cx, y + cy)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", recieveCommand)
	log.Fatal(http.ListenAndServe(*addr, nil))
}