package main

import (
	"./hashs"
	"./words"
	"math"
	"time"
)

const hashTobench = 10 * 1000 * 1000

func BenchHasher() int {

	var start = time.Now().UnixNano()
	var hasher = hashs.NewHasher()

	for i := 0; i < hashTobench; i++ {
		hasher.Hash("1234567890")
	}

	var end = time.Now().UnixNano()
	var timeInSeconds = float64(end-start) / 1000000000

	return int(math.Floor(hashTobench / timeInSeconds))
}

func BenchBruter() int {

	var start = time.Now().UnixNano()

	var worder = words.NewWorder([]string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"})

	for i := 0; i < hashTobench; i++ {
		worder.Next()
	}

	var end = time.Now().UnixNano()
	var timeInSeconds = float64(end-start) / 1000000000

	return int(math.Floor(hashTobench / timeInSeconds))
}
