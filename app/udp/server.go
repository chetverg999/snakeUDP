package udp

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"snacke/app/game"
	"snacke/app/game/config"
	"snacke/app/game/entity"
	"time"
)

var app = game.Game{}
var setting = game.Setting
var udpMessagesClient = make(chan string)

func Server() {
	connection, clientAddress := StartServer()

	if connection == nil {
		fmt.Println("Ошибка: сервер не запустился")
		return
	} // запуск сервера

	keysEvents, _ := keyboard.GetKeys(10)
	defer func() {
		_ = keyboard.Close()
	}()

	tickerSnake := time.NewTicker(game.Setting.GetDuration())
	tickerFood := time.NewTicker(game.Setting.GetFoodLiveDuration())
	defer tickerSnake.Stop()
	defer tickerFood.Stop()

	go readUDP(connection, udpMessagesClient)

	app.Init(StartPoint) // запуск игры

	for {
		select {

		case <-tickerFood.C:
			app.TickerDrawFood()
			break
		case <-tickerSnake.C:
			if isEaten := app.TickerMove(); isEaten {
				tickerFood = time.NewTicker(game.Setting.GetFoodLiveDuration())
			}
			break

		case event := <-keysEvents:

			switch event.Key {

			case keyboard.KeyArrowUp:
				app.Snakes()[0].ChangeDirection(entity.MoveUp())
				writeUDP(clientAddress, []byte(string(config.KeyArrowUp)), connection)
				break
			case keyboard.KeyArrowDown:
				app.Snakes()[0].ChangeDirection(entity.MoveDown())
				writeUDP(clientAddress, []byte(string(config.KeyArrowDown)), connection)
				break
			case keyboard.KeyArrowLeft:
				app.Snakes()[0].ChangeDirection(entity.MoveLeft())
				writeUDP(clientAddress, []byte(string(config.KeyArrowLeft)), connection)
				break
			case keyboard.KeyArrowRight:
				app.Snakes()[0].ChangeDirection(entity.MoveRight())
				writeUDP(clientAddress, []byte(string(config.KeyArrowRight)), connection)
				break
			case keyboard.KeyEsc:
				panic("Stop")
			default:
				continue
			}
		case message := <-udpMessagesClient:

			switch message {

			case string(config.MultiArrowUp):
				app.Snakes()[len(app.Snakes())-1].ChangeDirection(entity.MoveUp())
			case string(config.MultiArrowDown):
				app.Snakes()[len(app.Snakes())-1].ChangeDirection(entity.MoveDown())
			case string(config.MultiArrowLeft):
				app.Snakes()[len(app.Snakes())-1].ChangeDirection(entity.MoveLeft())
			case string(config.MultiArrowRight):
				app.Snakes()[len(app.Snakes())-1].ChangeDirection(entity.MoveRight())
			default:
				continue
			}
		}
	}
}
