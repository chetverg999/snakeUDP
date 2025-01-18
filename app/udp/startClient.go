package udp

import (
	"fmt"
	"net"
	"strconv"
)

var StartPoint int64

func startClient() (*net.UDPConn, *net.UDPAddr) {
	fmt.Println("Вы запустили клиентское приложение")
	serverAddress, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	connection, err := net.DialUDP("udp", nil, serverAddress)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	//defer connection.Close()

	// отправляем сообщение серверу
	_, err = connection.Write([]byte("Соединение установлено"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Соединение установлено")

	for {
		inputBytes := make([]byte, 50)
		fmt.Println("Ожидание стартовой точки от сервера...")
		n, serverAddress, err := connection.ReadFromUDP(inputBytes)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Received message from", serverAddress)
		fmt.Println(string(inputBytes[:n]))
		StartPoint, _ = strconv.ParseInt(string(inputBytes[:n]), 10, 64)
		break
	}

	return connection, serverAddress
}
