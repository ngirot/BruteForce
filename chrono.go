package main

import (
	"time"
)

type Chrono struct {
	start int64
	end int64
	state uint8
}

func NewChrono() Chrono {
	return Chrono{state:0}
}

func (c *Chrono) Start() {
	c.start = now()
	c.state = 1
}

func (c *Chrono) End() {
	c.end = now()
	c.state = 2
}

func (c *Chrono) DurationInNano() int64 {
	if c.state == 1 {
		return (now() - c.start)
	}
	if c.state == 2 {
		return c.end - c.start
	}

	return 0
}

func (c *Chrono) DurationInMilli() float64 {
	return float64(c.DurationInNano())/1000000
}

func (c *Chrono) DurationInSeconds() float64 {
	return c.DurationInMilli()/1000
}

func now() int64 {
	return time.Now().UnixNano()
}


