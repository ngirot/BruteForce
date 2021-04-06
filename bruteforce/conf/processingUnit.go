package conf

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

func (conf *ProcessingUnitConfiguration) Type() ProcessingUnit {
	return conf.unitType
}

type ProcessingUnit int

const (
	Cpu ProcessingUnit = iota
	Gpu
)
