package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)



func Reader(conn *websocket.Conn){
	for{
		messageType, message, err := conn.ReadMessage()
		if err != nil{
			log.Println(err.Error())
			return
		}

		log.Println(string(message))

		if err := conn.WriteMessage(messageType, message); err != nil{
			log.Println(err.Error())
			return
		}

	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is HTTP")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}

	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		err= errors.Wrap(err, "fail upgrade")
		log.Println(err.Error())
	}
	log.Println("Websocket connection success")
	Reader(websocket)
}