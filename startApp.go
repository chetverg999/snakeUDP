package main

import (
	"fmt"
	"snacke/app/game/config"
	"snacke/app/udp"
)

func main() {
	fmt.Print(config.AsciiArt)
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
			udp.Server()
			return
		case "C":
			udp.Client()
			return
		default:
			fmt.Print("Введено неизвестное значение, попробуйте снова: ")
		}
	}
}
