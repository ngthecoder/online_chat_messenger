package main

import (
	"fmt"
	"net"
)

func main() {
	//TCP server (chatroom related)
	tcpServerAddr := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
		Zone: "",
	}
	tcpListener, err := net.ListenTCP("tcp", tcpServerAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer tcpListener.Close()
	fmt.Println("TCP server listening on :8080")
	for {
		tcpConn, err := tcpListener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		go handleTCPConn(tcpConn)
	}
}

func handleTCPConn(tcpConn net.Conn) {
	buff := make([]byte, 1024)
	_, err := tcpConn.Read(buff)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	roomNameSize := buff[0]
	operation := buff[1]
	state := buff[2]
	loadSize := buff[3:32]
	roomname := buff[32 : 32+roomNameSize]
	payload := buff[32+roomNameSize:]

	fmt.Println(roomNameSize)
	fmt.Println(operation)
	fmt.Println(state)
	fmt.Println(loadSize)
	fmt.Println(roomname)
	fmt.Println(payload)

	responce := []byte{}
	responce = append(responce, operation)
	responce = append(responce, state)
	responce = append(responce, payload...)
	tcpConn.Write(responce)
}

func handleUDPConn() {
	// UDP server (messaging)
	udpServerAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8090,
		Zone: "",
	}
	messageConn, err := net.ListenUDP("udp", udpServerAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer messageConn.Close()
	fmt.Println("UDP server listening on :8090")

	users := make(map[string]User)

	for {
		buff := make([]byte, 1024)
		_, clientAddr, err := messageConn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		user := User{
			Usernamelen: int(buff[0]),
			Username:    string(buff[1 : 1+int(buff[0])]),
			IP:          clientAddr,
		}
		if _, exist := users[user.Username]; !exist {
			users[user.Username] = user
		}

		message := Message{
			Username: string(buff[1 : 1+int(buff[0])]),
			Content:  string(buff[1+int(buff[0]):]),
		}

		fmt.Println(user)
		fmt.Println(message)
		for _, individualUser := range users {
			messageConn.WriteToUDP([]byte(message.Username+": "+message.Content), individualUser.IP)
		}
	}
}
