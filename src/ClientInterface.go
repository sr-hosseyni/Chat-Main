package main

type ClientInterface interface {
	read()
	write()
	listen()
	outbound(message Message)
	inbound() Message
}