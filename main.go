package main

import (
	"fmt"
	"net"
)

type User struct {
	Usernamelen int
	Username    string
	IP          *net.UDPAddr
}

type Message struct {
	Username string
	Content  string
}

func main() {
	serverAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
		Zone: "",
	}
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("UDP server listening on :8080")

	for {
		buff := make([]byte, 1024)
		_, clientAddr, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		user := User{
			Usernamelen: int(buff[0]),
			Username:    string(buff[1 : 1+int(buff[0])]),
			IP:          clientAddr,
		}
		message := Message{
			Username: string(buff[1 : 1+int(buff[0])]),
			Content:  string(buff[1+int(buff[0]):]),
		}
		fmt.Println(user)
		fmt.Println(message)
		conn.WriteToUDP([]byte(message.Content), clientAddr)
	}
}
