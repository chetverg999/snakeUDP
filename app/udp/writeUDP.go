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
