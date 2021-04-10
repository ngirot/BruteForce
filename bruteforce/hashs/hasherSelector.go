package hashs

import (
	"errors"
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"github.com/ngirot/BruteForce/bruteforce/hashs/selector"
	"sort"
	"strings"
)

func HasherCreator(hashType string, processingUnitConfiguration conf.ProcessingUnitConfiguration) (func() hashers.Hasher, error) {
	var hasherMap, err = buildHasherMap(processingUnitConfiguration.Type())
	if err != nil {
		return nil, err
	}

	if description, present := hasherMap[hashType]; present {
		var creator = func() hashers.Hasher {
			return description.Build()
		}
		return creator, nil
	}

	return nil, errors.New("'" + hashType + "' is not a valid hash type, must be one of: " + strings.Join(AllHasherTypes(), ", "))
}

func HasherBenchmarkDescription(hashType string) string {
	var hasherMap, _ = buildHasherMap(conf.Cpu)
	if description, present := hasherMap[hashType]; present {
		return description.Description()
	}

	return hashType
}

func AllHasherTypes() []string {
	var values []string

	hasherMap, _ := buildHasherMap(conf.Cpu)
	for k := range hasherMap {
		values = append(values, k)
	}

	sort.Strings(values)

	return values
}

func IsValidHash(hash conf.HashConf) bool {
	if hasherCreator, e := HasherCreator(hash.HashType, conf.NewProcessingUnitConfiguration(false)); e == nil {
		return hasherCreator().IsValid(hash.Value)
	} else {
		return true
	}
}

func ExampleHash(hash conf.HashConf) string {
	if hasherCreator, e := HasherCreator(hash.HashType, conf.NewProcessingUnitConfiguration(false)); e == nil {
		return hasherCreator().Example()
	} else {
		return ""
	}
}

func buildHasherMap(processingUnit conf.ProcessingUnit) (map[string]selector.HasherInfos, error) {
	if processingUnit == conf.Gpu {
		return selector.BuildGpuHasherMap()
	} else {
		return selector.BuildCpuHasherMap(), nil
	}
}
