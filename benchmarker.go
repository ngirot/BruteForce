package main

import (
	"time"
	"math"
)

const hashTobench = 10*1000*1000

func BenchHasher() int {

	var start = time.Now().UnixNano()
	var hasher = NewHasher()

	for i:=0; i<hashTobench ; i++  {
		hasher.Hash("1234567890")
	}

	var end = time.Now().UnixNano()
	var timeInSeconds = float64(end - start) / 1000000000

	return int(math.Floor(hashTobench / timeInSeconds))
}

func BenchBruter() int {

	var start = time.Now().UnixNano()

	var worder = NewWorder()

	for i:=0; i<hashTobench ; i++  {
		worder.Next()
	}

	var end = time.Now().UnixNano()
	var timeInSeconds = float64(end - start) / 1000000000

	return int(math.Floor(hashTobench / timeInSeconds))
}



