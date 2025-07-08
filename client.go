package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

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

func main() {
	// TCP server for chatroom
	tcpServerAddr := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
		Zone: "",
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpServerAddr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer tcpConn.Close()

	fmt.Printf("Do you wish to create a new chatroom (Y/N): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	newChatroomRequest := scanner.Text()
	var chatroomName string
	if newChatroomRequest == "Y" || newChatroomRequest == "y" {
		fmt.Printf("Enter a name for new chatroom: ")
		scanner.Scan()
		chatroomName = scanner.Text()
	} else {
		fmt.Printf("Enter which chatroom you wish to join: ")
		scanner.Scan()
		chatroomName = scanner.Text()
	}

	// UDP server for messaging
	udpServerAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8090,
		Zone: "",
	}
	udpConn, err := net.DialUDP("udp", nil, udpServerAddr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer udpConn.Close()

	fmt.Printf("Enter username: ")
	scanner.Scan()
	username := scanner.Text()

	for {
		fmt.Printf("Enter message: ")
		scanner.Scan()
		message := scanner.Text()

		load := []byte{}
		load = append(load, byte(len(username)))
		load = append(load, []byte(username)...)
		load = append(load, []byte(message)...)

		udpConn.Write(load)

		buff := make([]byte, 1024)
		n, err := udpConn.Read(buff)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("Received: %s\n", string(buff[:n]))
	}
}
