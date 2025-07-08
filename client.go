package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddress := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
		Zone: "",
	}
	conn, err := net.DialUDP("udp", nil, serverAddress)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
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

		conn.Write(load)

		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("Received %s\n", string(buff[:n]))
	}
}
