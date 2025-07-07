package main

import (
	"fmt"
	"net"
)

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
		n, clientAddr, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Received %s from %s", string(buff[:n]), clientAddr.IP)
		conn.WriteToUDP(buff[:n], clientAddr)
	}
}
