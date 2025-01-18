package entity

import (
	"math/rand"
	"snacke/app/game/setting"
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

func NewFood(setting setting.Settings) *Food {
	return &Food{
		point: Point{
			X: rand.Intn(setting.Width-2) + 1,
			Y: rand.Intn(setting.Height-2) + 1},
	}
}
