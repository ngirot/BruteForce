package conf

import (
	"errors"
	"math"
)

type ProcessingUnitConfiguration struct {
	unitType ProcessingUnit
}

func NewProcessingUnitConfiguration(gpu bool) ProcessingUnitConfiguration {
	if gpu {
		return ProcessingUnitConfiguration{Gpu}
	} else {
		return ProcessingUnitConfiguration{Cpu}
	}
}

func (conf *ProcessingUnitConfiguration) NumberOfGoRoutines() int {
	if conf.unitType == Gpu {
		return 1
	} else {
		return BestNumberOfGoRoutine()
	}
}

func (conf *ProcessingUnitConfiguration) NumberOfWordsPerIteration() int {
	if conf.unitType == Gpu {
		return 100000
	} else {
		return 1
	}
}

func (conf *ProcessingUnitConfiguration) NumberOfWildcardsForDeportedComputingUnit(charsetLength int) int {
	if conf.unitType == Gpu {
		var i = 2
		for {
			if math.Pow(float64(charsetLength), float64(i)) > 750000 {
				return i - 1
			}
			i++
		}
	} else {
		return 0
	}
}

func (conf *ProcessingUnitConfiguration) CheckAvailability() error {
	if conf.unitType == Gpu && !HasDeviceAvailable() {
		return errors.New("No GPU found (when searching device with OpenCl compatibility)")
	} else {
		return nil
	}
}

func (conf *ProcessingUnitConfiguration) Type() ProcessingUnit {
	return conf.unitType
}

type ProcessingUnit int

const (
	Cpu ProcessingUnit = iota
	Gpu
)
