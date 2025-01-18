package udp

import (
	"fmt"
	"net"
	"time"
)

var addr = "192.168.0.104:8080"
var localAddr = "127.0.0.1:8080"
var Now = time.Now().UnixNano()

func StartServer() (*net.UDPConn, *net.UDPAddr) {
	fmt.Println("Вы запустили серверное приложение")
	serverAddress, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	connection, err := net.ListenUDP("udp", serverAddress)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	//defer connection.Close()

	// получаем сообщение от клиента
	inputBytes := make([]byte, 50)
	fmt.Println("Ожидание ответа от клиента...")
	n, clientAddress, err := connection.ReadFromUDP(inputBytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Received message from", clientAddress)
	fmt.Println(string(inputBytes[:n]))

	writeUDP(clientAddress, []byte(string(Now)), connection)
	fmt.Println("Отправлена стартовая точка для игры")

	return connection, clientAddress
}
