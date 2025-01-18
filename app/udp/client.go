package udp

import (
	"github.com/eiannone/keyboard"
	"snacke/app/game/config"
	"snacke/app/game/entity"
	"time"
)

var udpMessagesServer = make(chan string)

func Client() {
	connection, _ := startClient()

	keysEvents, _ := keyboard.GetKeys(10)
	defer func() {
		_ = keyboard.Close()
	}()

	tickerSnake := time.NewTicker(setting.GetDuration())
	tickerFood := time.NewTicker(setting.GetFoodLiveDuration())
	defer tickerSnake.Stop()
	defer tickerFood.Stop()

	go readUDP(connection, udpMessagesServer)

	app.Init(StartPoint)

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
		case message := <-udpMessagesServer:
			// получаем сообщение от сервера
			switch message {

			case string(config.KeyArrowUp):
				app.Snakes()[0].ChangeDirection(entity.MoveUp())
				break
			case string(config.KeyArrowDown):
				app.Snakes()[0].ChangeDirection(entity.MoveDown())
				break
			case string(config.KeyArrowLeft):
				app.Snakes()[0].ChangeDirection(entity.MoveLeft())
				break
			case string(config.KeyArrowRight):
				app.Snakes()[0].ChangeDirection(entity.MoveRight())
				break
			}

		case event := <-keysEvents:

			switch event.Key {

			default: // отправляем сообщение

				if entity.CheckMultiDirection(event.Rune, config.MultiArrowUp) {
					app.Snakes()[len(app.Snakes())-1].ChangeDirection(entity.MoveUp())
					writeUDPdial([]byte(string(config.MultiArrowUp)), connection)
				}

				if entity.CheckMultiDirection(event.Rune, config.MultiArrowDown) {
					app.Snakes()[len(app.Snakes())-1].ChangeDirection(entity.MoveDown())
					writeUDPdial([]byte(string(config.MultiArrowDown)), connection)
				}

				if entity.CheckMultiDirection(event.Rune, config.MultiArrowLeft) {
					app.Snakes()[len(app.Snakes())-1].ChangeDirection(entity.MoveLeft())
					writeUDPdial([]byte(string(config.MultiArrowLeft)), connection)
				}

				if entity.CheckMultiDirection(event.Rune, config.MultiArrowRight) {
					app.Snakes()[len(app.Snakes())-1].ChangeDirection(entity.MoveRight())
					writeUDPdial([]byte(string(config.MultiArrowRight)), connection)
				}
			}
		}
	}
}
