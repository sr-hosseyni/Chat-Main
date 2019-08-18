package main

import (
	"net"
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"bufio"
	"encoding/json"
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	//db := getDB()
	//defer db.Close()
	//// Migrate the schema
	//db.AutoMigrate(&Product{})

	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	chatRoom := getChatRoom()

	log.Println("net server started on :6666")
	listener, _ := net.Listen("tcp", ":6666")


	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		ws, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Fatal(err)
		}
		chatRoom.wsJoins <- *ws
	})

	http.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		w := bufio.NewWriter(writer)

		onlineUsers := UserList{
			UserList: make([]IUser, len(chatRoom.clients)),
		}

		i := 0
		for userId := range chatRoom.clients {
			log.Println(userId, users[userId],)
			onlineUsers.UserList[i] = users[userId].transform()
			i++
		}

		log.Println(onlineUsers)
		rsp, err := json.Marshal(onlineUsers)
		if err != nil {
			log.Println(err)
		}
		w.Write(rsp)
		w.Flush()
	})



	go func() {
			log.Println("http server started on :7777")
			err := http.ListenAndServe(":7777", nil)
			if err != nil {
				log.Fatal("ListenAndServe: ", err)
			}
	} ()

	for {
		conn, _ := listener.Accept()
		log.Println("new connection")
		chatRoom.tcpJoins <- conn
	}
}
