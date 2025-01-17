package game

import "time"

type Settings struct {
	Height           int
	Width            int
	PlayerCount      int
	Level            int
	FoodLiveDuration time.Duration
	duration         time.Duration
}

func (s *Settings) GetDuration() time.Duration {
	index := 500 / s.Level

	return time.Duration(index) * time.Millisecond
}

func (s *Settings) GetFoodLiveDuration() time.Duration {
	return s.FoodLiveDuration
}

var Setting = Settings{
	Height:           20,
	Width:            40,
	PlayerCount:      2,
	FoodLiveDuration: 5 * time.Second,
	Level:            5,
}
