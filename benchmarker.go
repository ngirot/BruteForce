package main

import (
	"./hashs"
	"./words"
	"math"
)

const hashTobench = 10 * 1000 * 1000

func BenchHasher() int {
	var hasher = hashs.NewHasher()

	var chrono = NewChrono()
	chrono.Start()

	for i := 0; i < hashTobench; i++ {
		hasher.Hash("1234567890")
	}

	chrono.End()

	return int(math.Floor(hashTobench / chrono.DurationInSeconds()))
}

func BenchBruter() int {
	var worder = words.NewWorder([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, 1, 0)

	var chrono = NewChrono()
	chrono.Start()

	for i := 0; i < hashTobench; i++ {
		worder.Next()
	}

	chrono.End()

	return int(math.Floor(hashTobench / chrono.DurationInSeconds()))
}
