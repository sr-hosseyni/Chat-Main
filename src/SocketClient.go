package main

import (
	"bufio"
	"net"
	"encoding/json"
	"log"
	"fmt"
)

type SocketClient struct {
	incoming 	chan Message
	outgoing 	chan Message
	jsonReader  *json.Decoder
	jsonWriter  *json.Encoder
	writer 		*bufio.Writer
	reader 		*bufio.Reader
}

func (client *SocketClient) read() {

	var user User
	err := client.jsonReader.Decode(&user)
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
		err := client.jsonReader.Decode(&msg)
		fmt.Println(msg)
		if err != nil {
			log.Printf("error: %v", err)
			fmt.Printf("error: %v", err)
			break
		}
		msg.From = users[usersNameIndex[user.Name]].Id
		client.incoming <- msg
	}
}

func (client *SocketClient) write() {
	for data := range client.outgoing {
		log.Println("sending msg to client")
		client.jsonWriter.Encode(data)
		client.writer.Flush()
	}
}

func (client *SocketClient) listen() {
	go client.read()
	go client.write()
}

func (client *SocketClient) outbound(message Message) {
	client.outgoing <- message
}

func (client *SocketClient) inbound() Message {
	return <- client.incoming
}

func newSocketClient(connection net.Conn) *SocketClient {
	fmt.Println("new connection")
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &SocketClient {
		incoming: make(chan Message),
		outgoing: make(chan Message),
		jsonReader: json.NewDecoder(reader),
		jsonWriter: json.NewEncoder(writer),
		writer: writer,
		reader: reader,
	}

	client.listen()

	return client
}
