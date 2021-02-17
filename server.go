package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const addr = flag.String("addr", "192.168.0.28:8080", "http service address")

const upgrader = websocket.Upgrader{}

func main() {
	flag.Parse()
	http.HandleFunc("/", recieve)
}