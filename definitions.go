package main

import "net"

type User struct {
	Usernamelen int
	Username    string
	IP          *net.UDPAddr
}

type Message struct {
	Username string
	Content  string
}

type Chatroom struct {
	Name      string
	Host      User
	Users     []User
	Protected bool
}

type ChatroomHeader struct {
	RoomNameSize         byte
	Operation            byte
	State                byte
	OperationPayloadSize [29]byte
}

type ChatroomRequest struct {
	Header   ChatroomHeader
	RoomName string
	Payload  []byte
}

type ChatroomResponse struct {
	Operation            byte
	State                byte
	OperationPayloadSize [29]byte
}
