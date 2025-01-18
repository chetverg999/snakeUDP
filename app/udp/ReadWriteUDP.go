package udp

import (
	"fmt"
	"net"
)

func writeUDPdial(message []byte, conn *net.UDPConn) {
	_, err := conn.Write(message)
	if err != nil {
		fmt.Println(err)
	}
}

func writeUDP(address *net.UDPAddr, message []byte, conn *net.UDPConn) {
	_, err := conn.WriteToUDP(message, address)
	if err != nil {
		fmt.Println(err)
	}
}

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
