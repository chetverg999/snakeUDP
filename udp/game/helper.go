package game

import "fmt"

type Point struct {
	X, Y int
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func MoveUp() Point {
	return Point{X: 0, Y: -1}
}

func MoveDown() Point {
	return Point{X: 0, Y: 1}
}

func MoveLeft() Point {
	return Point{X: -1, Y: 0}
}

func MoveRight() Point {
	return Point{X: 1, Y: 0}
}

func CheckMultiDirection(key rune, available []rune) bool {
	for _, valueAvailable := range available {
		if key == valueAvailable {
			return true
		}
	}

	return false
}
