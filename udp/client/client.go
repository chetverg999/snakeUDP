package client

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"net"
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
var inputBytes = make([]byte, 1)

func Client() {
	connection, serverAddress := startClient()
	app.Init() // запуск игры

	keysEvents, _ := keyboard.GetKeys(10)
	defer func() {
		_ = keyboard.Close()
	}()

	tickerSnake := time.NewTicker(setting.GetDuration())
	tickerFood := time.NewTicker(setting.GetFoodLiveDuration())
	defer tickerSnake.Stop()
	defer tickerFood.Stop()

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
				break
			case keyboard.KeyArrowDown:
				app.Snakes()[0].ChangeDirection(game.MoveDown())
				break
			case keyboard.KeyArrowLeft:
				app.Snakes()[0].ChangeDirection(game.MoveLeft())
				break
			case keyboard.KeyArrowRight:
				app.Snakes()[0].ChangeDirection(game.MoveRight())
				break
			case keyboard.KeyEsc:
				panic("Stop")
			default: // получаем сообщение от клиента

				if game.CheckMultiDirection(event.Rune, game.MultiArrowUp) {
					_, err := connection.Write([]byte(string(game.MultiArrowUp)))
					if err != nil {
						fmt.Println(err)
					}
				}

				if game.CheckMultiDirection(event.Rune, game.MultiArrowDown) {
					_, err := connection.Write([]byte(string(game.MultiArrowDown)))
					if err != nil {
						fmt.Println(err)
					}
				}

				if game.CheckMultiDirection(event.Rune, game.MultiArrowLeft) {
					_, err := connection.Write([]byte(string(game.MultiArrowLeft)))
					if err != nil {
						fmt.Println(err)
					}
				}

				if game.CheckMultiDirection(event.Rune, game.MultiArrowRight) {
					_, err := connection.Write([]byte(string(game.MultiArrowRight)))
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}

	_, err = connection.Read(inputBytes)
	if err != nil {
		fmt.Println(err)
		continue
	}
}
