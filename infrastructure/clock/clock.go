package clock

import "time"

type Clock interface {
	Now() time.Time
}

type TimeClock struct{}

func NewTimeClock() *TimeClock {
	return &TimeClock{}
}

func (c *TimeClock) Now() time.Time {
	return time.Now()
}
