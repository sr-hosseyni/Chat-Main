package main

import (
	"net"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
	"sync"
)

var (
	SingletonChatRoom   *ChatRoom
	once				sync.Once
)

func getChatRoom() *ChatRoom {
	once.Do(func() {
		SingletonChatRoom = newChatRoom()
	})

	return SingletonChatRoom
}

type ChatRoom struct {
	clients map[int]ClientInterface
	tcpJoins chan net.Conn
	wsJoins  chan websocket.Conn
	incoming chan Message
	outgoing chan Message
}

func (chatRoom *ChatRoom) Broadcast(data Message) {
	log.Println("broadcasting message", data, chatRoom.clients)
	for _, client := range chatRoom.clients {
		client.outbound(data)
	}
}

func (chatRoom *ChatRoom) Send(data Message) {
	log.Println("Sending message ", data)
	if chatRoom.clients[data.To] != nil {
		chatRoom.clients[data.To].outbound(data)
	} else {
		// @todo handle offline message
	}
}

func (chatRoom *ChatRoom) webJoin(connection websocket.Conn) {
	fmt.Println("new web join")
	client := newWebSocketClient(connection)
	//chatRoom.clients = append(chatRoom.clients, client)
	go func() { for { chatRoom.incoming <- <-client.incoming } }()
}

func (chatRoom *ChatRoom) tcpJoin(connection net.Conn) {
	fmt.Println("new tcp join")
	client := newSocketClient(connection)
	//chatRoom.clients = append(chatRoom.clients, client)
	go func() { for { chatRoom.incoming <- <-client.incoming } }()
}

func (chatRoom *ChatRoom) listen() {
	go func() {
		fmt.Println("chat room is listening")
		for {
			select {
			case data := <-chatRoom.incoming:
				//chatRoom.Broadcast(data)
				chatRoom.Send(data)
			case conn := <-chatRoom.tcpJoins:
				chatRoom.tcpJoin(conn)
			case conn := <-chatRoom.wsJoins:
				chatRoom.webJoin(conn)
			}
		}
	}()
}

func newChatRoom() *ChatRoom {
	chatRoom := &ChatRoom{
		clients:  make(map[int]ClientInterface, 0),
		tcpJoins: make(chan net.Conn),
		wsJoins:  make(chan websocket.Conn),
		incoming: make(chan Message),
		outgoing: make(chan Message),
	}

	chatRoom.listen()

	return chatRoom
}
