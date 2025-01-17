package server

import (
	"fmt"
	"net"
)

func StartServer() (*net.UDPConn, *net.UDPAddr) {
	fmt.Println("Вы запустили серверное приложение")
	serverAddress, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	connection, err := net.ListenUDP("udp", serverAddress)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer connection.Close()

	// получаем сообщение от клиента
	inputBytes := make([]byte, 50)
	fmt.Println("Ожидание ответа от клиента...")
	n, clientAddress, err := connection.ReadFromUDP(inputBytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Received message from", clientAddress)
	fmt.Println(string(inputBytes[:n]))

	return connection, clientAddress
}
