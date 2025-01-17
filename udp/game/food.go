package game

import (
	"math/rand"
)

type Food struct {
	point Point
}

func (f *Food) Point() Point {
	return f.point
}

func (f *Food) SetPoint(point Point) {
	f.point = point
}

func NewFood() *Food {
	return &Food{
		point: Point{
			X: rand.Intn(Setting.Width-2) + 1,
			Y: rand.Intn(Setting.Height-2) + 1},
	}
}
