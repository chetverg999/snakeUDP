package server

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"snacke/udp/game"
	"time"
)

const (
	KeyArrowUp = 0 + iota
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
)

var app = game.Game{}
var setting = game.Setting
var udpMessages = make(chan string)

func Server() {
	connection, clientAddress := StartServer()
	if connection == nil {
		fmt.Println("Ошибка: сервер не запустился")
		return
	} // запуск сервера
	app.Init() // запуск игры

	keysEvents, _ := keyboard.GetKeys(10)
	defer func() {
		_ = keyboard.Close()
	}()

	tickerSnake := time.NewTicker(setting.GetDuration())
	tickerFood := time.NewTicker(setting.GetFoodLiveDuration())
	defer tickerSnake.Stop()
	defer tickerFood.Stop()

	go func() {
		inputBytes := make([]byte, 1024)
		for {
			n, _, err := connection.ReadFromUDP(inputBytes)
			if err != nil {
				fmt.Println("Ошибка при чтении из UDP:", err)
				continue
			}

			message := string(inputBytes[:n])
			if len(message) > 0 { // Игнорируем пустые сообщения
				udpMessages <- message
			}
		}
	}()

	for {
		select {

		case <-tickerFood.C:
			app.TickerDrawFood()
			break
		case <-tickerSnake.C:
			if isEaten := app.TickerMove(); isEaten {
				tickerFood = time.NewTicker(setting.GetFoodLiveDuration())
			}
			break

		case event := <-keysEvents:

			switch event.Key {

			case keyboard.KeyArrowUp:
				app.Snakes()[0].ChangeDirection(game.MoveUp())
				// отправляем сообщение клиенту
				_, err := connection.WriteToUDP([]byte(string(KeyArrowUp)), clientAddress)
				if err != nil {
					fmt.Println(err)
					continue
				}
				break
			case keyboard.KeyArrowDown:
				app.Snakes()[0].ChangeDirection(game.MoveDown())
				// отправляем сообщение клиенту
				_, err := connection.WriteToUDP([]byte(string(KeyArrowDown)), clientAddress)
				if err != nil {
					fmt.Println(err)
					continue
				}
				break
			case keyboard.KeyArrowLeft:
				app.Snakes()[0].ChangeDirection(game.MoveLeft())
				// отправляем сообщение клиенту
				_, err := connection.WriteToUDP([]byte(string(KeyArrowLeft)), clientAddress)
				if err != nil {
					fmt.Println(err)
					continue
				}
				break
			case keyboard.KeyArrowRight:
				app.Snakes()[0].ChangeDirection(game.MoveRight())
				// отправляем сообщение клиенту
				_, err := connection.WriteToUDP([]byte(string(KeyArrowRight)), clientAddress)
				if err != nil {
					fmt.Println(err)
					continue
				}
				break
			case keyboard.KeyEsc:
				panic("Stop")
			default:
				continue
			}
		case messege := <-udpMessages:

			switch messege {

			case string(game.MultiArrowUp):
				app.Snakes()[len(app.Snakes())-1].ChangeDirection(game.MoveUp())
			case string(game.MultiArrowDown):
				app.Snakes()[len(app.Snakes())-1].ChangeDirection(game.MoveDown())
			case string(game.MultiArrowLeft):
				app.Snakes()[len(app.Snakes())-1].ChangeDirection(game.MoveLeft())
			case string(game.MultiArrowRight):
				app.Snakes()[len(app.Snakes())-1].ChangeDirection(game.MoveRight())
			default:
				continue
			}
		}
	}
}
