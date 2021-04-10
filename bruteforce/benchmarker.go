package bruteforce

import (
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"github.com/ngirot/BruteForce/bruteforce/words"
	"math"
	"time"
)

func BenchHasherOneCpu(hasherCreator func() hashers.Hasher) int {
	var buildActionFunc = getBuildActionFuncForHasher(hasherCreator)
	return bench(buildActionFunc, 1)
}

func BenchHasherMultiCpu(hasherCreator func() hashers.Hasher) int {
	var buildActionFunc = getBuildActionFuncForHasher(hasherCreator)
	return bench(buildActionFunc, conf.BestNumberOfGoRoutine())
}

func BenchWorderOneCpu() int {
	var buildActionFunc = getBuildActionFuncForWorder()
	return bench(buildActionFunc, 1)
}

func BenchWorderMultiCpu() int {
	var buildActionFunc = getBuildActionFuncForWorder()
	return bench(buildActionFunc, conf.BestNumberOfGoRoutine())
}

func bench(buildActionFunc func() func(), cpus int) int {
	var preBuiltActionFunc = make([]func(), cpus)
	for i := 0; i < cpus; i++ {
		preBuiltActionFunc[i] = buildActionFunc()
	}

	var chrono = NewChrono()
	chrono.Start()

	var count = 0
	var oneDone = func() {
		count++
	}

	var quit = make(chan bool)
	for i := 0; i < cpus; i++ {
		go actionLoop(preBuiltActionFunc[i], oneDone, quit)
	}

	time.Sleep(time.Second * 5)

	chrono.End()
	for i := 0; i < cpus; i++ {
		quit <- true
	}

	return int(math.Floor(float64(count) / chrono.DurationInSeconds()))
}

func actionLoop(action func(), oneDone func(), quit chan bool) {
	for {
		action()

		select {
		case <-quit:
			return
		default:
			oneDone()
		}
	}
}
func getBuildActionFuncForHasher(hasherCreator func() hashers.Hasher) func() func() {
	return func() func() {
		var hasher = hasherCreator()
		var referenceData = hasher.Example()
		return func() {
			hasher.Compare(hasher.Hash([]string{"1234567890"})[0], hasher.DecodeInput(referenceData))
		}
	}
}

func getBuildActionFuncForWorder() func() func() {
	return func() func() {
		var alphabet = words.DefaultAlphabet()
		var worder = words.NewWorderAlphabet(alphabet, 1, 0)
		return func() {
			worder.Next()
		}
	}
}
