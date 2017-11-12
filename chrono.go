package main

import (
	"time"
)

type Chrono interface {
	Start()
	End()
	DurationInNano() int64
	DurationInMilli() float64
	DurationInSeconds() float64
}

type chrono struct {
	start int64
	end   int64
	state uint8
}

func NewChrono() Chrono {
	return &chrono{state: 0}
}

func (c *chrono) Start() {
	c.start = now()
	c.state = 1
}

func (c *chrono) End() {
	c.end = now()
	c.state = 2
}

func (c *chrono) DurationInNano() int64 {
	if c.state == 1 {
		return (now() - c.start)
	}
	if c.state == 2 {
		return c.end - c.start
	}

	return 0
}

func (c *chrono) DurationInMilli() float64 {
	return float64(c.DurationInNano()) / 1000000
}

func (c *chrono) DurationInSeconds() float64 {
	return c.DurationInMilli() / 1000
}

func now() int64 {
	return time.Now().UnixNano()
}
