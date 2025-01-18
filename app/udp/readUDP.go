package udp

import (
	"fmt"
	"net"
)

func readUDP(connection *net.UDPConn, udpMessages chan string) {
	buffer := make([]byte, 1024)
	for {
		n, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении из UDP:", err)
			continue
		}
		message := string(buffer[:n])
		if message != "" {
			udpMessages <- message // Отправляем сообщение в канал
		}
	}
}
