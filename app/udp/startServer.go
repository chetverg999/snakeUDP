package udp

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var Now = time.Now().UnixNano()
var clientAddress *net.UDPAddr

func StartServer() (*net.UDPConn, *net.UDPAddr) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	addr, _ := os.LookupEnv("LOCAL_HOST")
	fmt.Println("Вы запустили серверное приложение")
	serverAddress, err := net.ResolveUDPAddr("udp", addr)
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
	secretKey := Secretkey()
	fmt.Println("Клиент должен отправить этот ключ:")
	fmt.Println(secretKey)

	// получаем сообщение от клиента
	inputBytes := make([]byte, 50)
	fmt.Println("Ожидание ответа от клиента...")

	for {
		n, clientAddress, err := connection.ReadFromUDP(inputBytes)
		if err != nil {
			fmt.Println(err)
		}
		if string(inputBytes[:n]) != strconv.FormatInt(secretKey, 10) {
			fmt.Println("Неправильный ключ")
			writeUDP(clientAddress, []byte("BAD"), connection)
			continue
		} else {
			writeUDP(clientAddress, []byte("OK"), connection)
			writeUDP(clientAddress, []byte(string(Now)), connection)
			fmt.Println("Отправлена стартовая точка для игры")
			break
		}
	}
	return connection, clientAddress
}
