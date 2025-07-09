package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

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

	request := []byte{}
	request = append(request, byte(len(chatroomName)))
	request = append(request, byte(0))
	request = append(request, byte(0))
	for i := 0; i < 29; i++ {
		request = append(request, byte(0))
	}
	request = append(request, []byte(chatroomName)...)
	tcpConn.Write([]byte(request))

	tcpResponse := make([]byte, 1024)
	_, err = tcpConn.Read(tcpResponse)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	operation := tcpResponse[0]
	state := tcpResponse[1]
	load := tcpResponse[2:]
	fmt.Println(operation)
	fmt.Println(state)
	fmt.Println(load)
}

func connectUDP() {
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
	scanner := bufio.NewScanner(os.Stdin)
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

		udpResponse := make([]byte, 1024)
		udpResponseSize, err := udpConn.Read(udpResponse)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("Received: %s\n", string(udpResponse[:udpResponseSize]))
	}
}
