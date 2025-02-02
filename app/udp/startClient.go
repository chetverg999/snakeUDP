package udp

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"strconv"
)

var StartPoint int64
var input string

func startClient() (*net.UDPConn, *net.UDPAddr) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	fmt.Println("Вы запустили клиентское приложение")
	addr, _ := os.LookupEnv("LOCAL_HOST")
	serverAddress, err := net.ResolveUDPAddr("udp", addr)
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
	inputBytes := make([]byte, 50)
	fmt.Println("Введите секретный ключ:")
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
		}
		writeUDPdial([]byte(input), connection)
		n, _, err := connection.ReadFromUDP(inputBytes)
		if string(inputBytes[:n]) != "OK" {
			fmt.Println("Неправильный ключ, попробуйте еще раз:")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Println("Ожидание стартовой точки от сервера...")
		n, _, err := connection.ReadFromUDP(inputBytes)
		if err != nil {
			fmt.Println(err)
		}
		StartPoint, _ = strconv.ParseInt(string(inputBytes[:n]), 10, 64)
		break
	}
	return connection, serverAddress
}
