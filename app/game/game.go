package game

import (
	"fmt"
	"math/rand"
	"snacke/app/game/config"
	"snacke/app/game/entity"
	"snacke/app/game/setting"
	"time"
)

var Setting = setting.Settings{
	Height:           10,
	Width:            20,
	PlayerCount:      2,
	FoodLiveDuration: 10 * time.Second,
	Level:            5,
}

type Game struct {
	snakes []*entity.Snake
	foods  []*entity.Food
	boards [][][]rune
	board  *entity.Board
	snake  *entity.Snake
	food   *entity.Food
}

func (g *Game) Board() *entity.Board {
	return g.board
}

func (g *Game) SetBoard(board *entity.Board) {
	g.board = board
}

func (g *Game) Snake() *entity.Snake {
	return g.snake
}

func (g *Game) SetSnake(snake *entity.Snake) {
	g.snake = snake
}

func (g *Game) Food() *entity.Food {
	return g.food
}

func (g *Game) SetFood(food *entity.Food) {
	g.food = food
}

func (g *Game) Boards() [][][]rune {
	return g.boards
}

func (g *Game) SetBoards(boards [][][]rune) {
	g.boards = boards
}

func (g *Game) Snakes() []*entity.Snake {
	return g.snakes
}

func (g *Game) SetSnakes(snakes []*entity.Snake) {
	g.snakes = snakes
}

func (g *Game) Foods() []*entity.Food {
	return g.foods
}

func (g *Game) SetFoods(foods []*entity.Food) {
	g.foods = foods
}

func (g *Game) Init(seed int64) {
	rand.Seed(seed)
	snakes := []*entity.Snake{entity.NewSnake(Setting)}
	foods := []*entity.Food{entity.NewFood(Setting)}

	if Setting.PlayerCount == config.MULTI {
		snakes = append(snakes, entity.NewSnake(Setting))
		foods = append(foods, entity.NewFood(Setting))
	}

	g.SetSnakes(snakes)
	g.SetFoods(foods)
	g.DrawBoard()
}

func (g *Game) DrawBoard() {
	entity.ClearScreen()
	g.SetBoards(make([][][]rune, 0))
	g.SetBoards(append(g.Boards(), g.Board().Create(Setting)))

	if Setting.PlayerCount == config.MULTI {
		g.SetBoards(append(g.Boards(), g.Board().Create(Setting)))
	}

	for numberSnake, snake := range g.Snakes() {
		for key, point := range snake.Body() {
			if len(snake.Body()) > 5 &&
				key != 0 &&
				snake.Body()[0].X == point.X &&
				snake.Body()[0].Y == point.Y {
				panic("BOOM")
			}

			g.Boards()[numberSnake][point.Y][point.X] = 'S'
		}
	}

	for numberFood, food := range g.Foods() {
		g.Boards()[numberFood][food.Point().Y][food.Point().X] = 'O'
	}

	for _, board := range g.Boards() {
		for i := 0; i < Setting.Height; i++ {
			for _, cell := range board[i] {
				fmt.Print(string(cell))
			}
			fmt.Println()
		}
		fmt.Print("\n\n\n\n")
	}
}

func (g *Game) TickerDrawFood() {
	g.SetFoods([]*entity.Food{})

	for _, _ = range g.Snakes() {
		g.SetFoods(append(g.Foods(), entity.NewFood(Setting)))
	}
}

func (g *Game) TickerMove() bool {
	var isEaten = false
	for number, snake := range g.Snakes() {
		snake.Move(Setting)

		if snake.Body()[0].X == g.Foods()[number].Point().X &&
			snake.Body()[0].Y == g.Foods()[number].Point().Y {
			snake.Grow()
			if number != len(g.Snakes())-1 {
				g.SetFoods(append([]*entity.Food{entity.NewFood(Setting)}, g.Foods()[number+1:]...))
			} else {
				g.SetFoods(append([]*entity.Food{entity.NewFood(Setting)}, g.Foods()[:len(g.Foods())-1]...))
			}
			isEaten = true
		}

	}
	g.DrawBoard()

	return isEaten
}
