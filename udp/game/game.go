package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	snakes []*Snake
	foods  []*Food
	boards [][][]rune
	board  *Board
	snake  *Snake
	food   *Food
}

func (g *Game) Board() *Board {
	return g.board
}

func (g *Game) SetBoard(board *Board) {
	g.board = board
}

func (g *Game) Snake() *Snake {
	return g.snake
}

func (g *Game) SetSnake(snake *Snake) {
	g.snake = snake
}

func (g *Game) Food() *Food {
	return g.food
}

func (g *Game) SetFood(food *Food) {
	g.food = food
}

func (g *Game) Boards() [][][]rune {
	return g.boards
}

func (g *Game) SetBoards(boards [][][]rune) {
	g.boards = boards
}

func (g *Game) Snakes() []*Snake {
	return g.snakes
}

func (g *Game) SetSnakes(snakes []*Snake) {
	g.snakes = snakes
}

func (g *Game) Foods() []*Food {
	return g.foods
}

func (g *Game) SetFoods(foods []*Food) {
	g.foods = foods
}

func (g *Game) Init() {
	rand.Seed(time.Now().UnixNano())
	snakes := []*Snake{NewSnake()}
	foods := []*Food{NewFood()}

	if Setting.PlayerCount == MULTI {
		snakes = append(snakes, NewSnake())
		foods = append(foods, NewFood())
	}

	g.SetSnakes(snakes)
	g.SetFoods(foods)
	g.DrawBoard()
}

func (g *Game) DrawBoard() {
	ClearScreen()
	g.SetBoards(make([][][]rune, 0))
	g.SetBoards(append(g.Boards(), g.Board().Create()))

	if Setting.PlayerCount == MULTI {
		g.SetBoards(append(g.Boards(), g.Board().Create()))
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
	g.SetFoods([]*Food{})

	for _, _ = range g.Snakes() {
		g.SetFoods(append(g.Foods(), NewFood()))
	}
}

func (g *Game) TickerMove() bool {
	var isEaten = false
	for number, snake := range g.Snakes() {
		snake.Move()

		if snake.Body()[0].X == g.Foods()[number].Point().X &&
			snake.Body()[0].Y == g.Foods()[number].Point().Y {
			snake.Grow()
			if number != len(g.Snakes())-1 {
				g.SetFoods(append([]*Food{NewFood()}, g.Foods()[number+1:]...))
			} else {
				g.SetFoods(append([]*Food{NewFood()}, g.Foods()[:len(g.Foods())-1]...))
			}
			isEaten = true
		}

	}
	g.DrawBoard()

	return isEaten
}
