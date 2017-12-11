package bruteforce

import (
	"math"
	"github.com/ngirot/BruteForce/bruteforce/hashs"
	"github.com/ngirot/BruteForce/bruteforce/words"
	"time"
	"github.com/ngirot/BruteForce/bruteforce/conf"
)

const hashTobench = 10 * 1000 * 1000

func BenchHasherOneCpu(hasherCreator func() hashs.Hasher) int {
	return bench(hasherCreator, 1)
}

func BenchHasherMultiCpu(hasherCreator func() hashs.Hasher) int {
	return bench(hasherCreator, conf.BestNumberOfGoRoutine())
}


func bench(hasherCreator func() hashs.Hasher, cpus int) int {
	var chrono = NewChrono()
	chrono.Start()

	var count = 0
	var oneDone = func() {
		count++
	}

	var quit = make(chan bool)
	for i := 0; i < cpus; i++ {
		go hashLoop(hasherCreator(), oneDone, quit)
	}


	time.Sleep(time.Second * 5)

	chrono.End()
	for i := 0; i < cpus; i++ {
		quit <- true
	}

	return int(math.Floor(float64(count) / chrono.DurationInSeconds()))
}

func hashLoop(hasher hashs.Hasher, oneDone func(), quit chan bool) {
	for {
		hasher.Hash("1234567890")

		select {
		case <- quit:
			return
		default:
			oneDone()
		}
	}
}

func BenchBruter() int {
	var alphabet = words.DefaultAlphabet()
	var worder = words.NewWorder(alphabet, 1, 0)

	var chrono = NewChrono()
	chrono.Start()

	for i := 0; i < hashTobench; i++ {
		worder.Next()
	}

	chrono.End()

	return int(math.Floor(hashTobench / chrono.DurationInSeconds()))
}
