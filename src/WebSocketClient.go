package main

import (
	"github.com/gorilla/websocket"
	"log"
	"fmt"
)

type WebSocketClient struct {
	incoming chan Message
	outgoing chan Message
	ws       *websocket.Conn
}

func (client *WebSocketClient) read() {
	defer client.ws.Close()


	var user User
	err := client.ws.ReadJSON(&user)
	if err != nil {
		log.Printf("ERRRRROR: %v", err)
		return
	}

	if !users[usersNameIndex[user.Name]].checkPassword(user.Password) {
		fmt.Println("Invalid user [" + user.Name + "]'s password")
		return
	}

	//client.user = users[user.Name]
	getChatRoom().clients[users[usersNameIndex[user.Name]].Id] = client

	fmt.Println(user.Name + " joined")


	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := client.ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		msg.From = users[usersNameIndex[user.Name]].Id
		// Send the newly received message to the broadcast channel
		client.incoming <- msg
	}
}

func (client *WebSocketClient) write() {
	for data := range client.outgoing {
		err := client.ws.WriteJSON(data)
		if err != nil {
			log.Printf("error: %v", err)
			client.ws.Close()
		}
	}
}

func (client *WebSocketClient) listen() {
	go client.read()
	go client.write()
}

func (client *WebSocketClient) outbound(message Message) {
	client.outgoing <- message
}

func (client *WebSocketClient) inbound() Message {
	return <- client.incoming
}

func newWebSocketClient(wsc websocket.Conn) *WebSocketClient {

	client := &WebSocketClient {
		incoming: make(chan Message),
		outgoing: make(chan Message),
		ws: &wsc,
	}

	client.listen()

	return client
}
