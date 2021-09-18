package main

import (
	"log"
	"net/http"

	"github.com/asianjack19/go-personal-projects/test-websocket/controller"
)


func setupRoute() {
	http.HandleFunc("/http", controller.HomePage)
	http.HandleFunc("/websocket", controller.WsEndpoint)
}

func main() {
	setupRoute()
	log.Fatal(http.ListenAndServe(":6969",nil))
}