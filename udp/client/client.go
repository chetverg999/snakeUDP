package client

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

func Client() {
	connection, _ := startClient()
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
		case messege := <-udpMessages:
			// получаем сообщение от сервера
			switch messege {

			case string(KeyArrowUp):
				app.Snakes()[0].ChangeDirection(game.MoveUp())
				break
			case string(KeyArrowDown):
				app.Snakes()[0].ChangeDirection(game.MoveDown())
				break
			case string(KeyArrowLeft):
				app.Snakes()[0].ChangeDirection(game.MoveLeft())
				break
			case string(KeyArrowRight):
				app.Snakes()[0].ChangeDirection(game.MoveRight())
				break
			}

		case event := <-keysEvents:

			switch event.Key {

			default: // отправляем сообщение

				if game.CheckMultiDirection(event.Rune, game.MultiArrowUp) {
					app.Snakes()[len(app.Snakes())-1].ChangeDirection(game.MoveUp())
					_, err := connection.Write([]byte(string(game.MultiArrowUp)))
					if err != nil {
						fmt.Println(err)
					}
				}

				if game.CheckMultiDirection(event.Rune, game.MultiArrowDown) {
					app.Snakes()[len(app.Snakes())-1].ChangeDirection(game.MoveDown())
					_, err := connection.Write([]byte(string(game.MultiArrowDown)))
					if err != nil {
						fmt.Println(err)
					}
				}

				if game.CheckMultiDirection(event.Rune, game.MultiArrowLeft) {
					app.Snakes()[len(app.Snakes())-1].ChangeDirection(game.MoveLeft())
					_, err := connection.Write([]byte(string(game.MultiArrowLeft)))
					if err != nil {
						fmt.Println(err)
					}
				}

				if game.CheckMultiDirection(event.Rune, game.MultiArrowRight) {
					app.Snakes()[len(app.Snakes())-1].ChangeDirection(game.MoveRight())
					_, err := connection.Write([]byte(string(game.MultiArrowRight)))
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}
