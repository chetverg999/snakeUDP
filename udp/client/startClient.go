package client

import (
	"fmt"
	"net"
)

func startClient() (*net.UDPConn, *net.UDPAddr) {
	fmt.Println("Вы запустили клиентское приложение")
	serverAddress, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	connection, err := net.DialUDP("udp", nil, serverAddress)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer connection.Close()

	_, err = connection.Write([]byte("Соединение установлено"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Соединение установлено")

	return connection, serverAddress
}
