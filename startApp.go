package main

import (
	"fmt"
	"snacke/udp/client"
	"snacke/udp/game"
	"snacke/udp/server"
)

func main() {
	fmt.Print(game.AsciiArt)
	fmt.Println("Вы запускаете сервер или клиент?")
	fmt.Println("Если сервер - введите S, если клиент - введите C")
	var input string
	fmt.Print("Введите значение запуска: ")
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		} // Читаем строку до нажатия Enter
		switch input {
		case "S":
			server.Server()
			return
		case "C":
			client.Client()
			return
		default:
			fmt.Print("Введено неизвестное значение, попробуйте снова: ")
		}
	}
}
