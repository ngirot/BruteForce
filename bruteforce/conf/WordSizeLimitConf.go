package conf

import (
	"fmt"
	"strconv"
	"strings"
)

type WorSizeLimitConf struct {
	Min int
	Max int
}

const rangeSeparator = "-"

func NewWordSizeLimitConf(param string) (WorSizeLimitConf, error) {
	if param == "" {
		return defaultConf(), nil
	}

	if strings.Contains(param, rangeSeparator) {
		return buildRange(param)
	} else {
		return buildSingleSize(param)
	}
}

func defaultConf() WorSizeLimitConf {
	return WorSizeLimitConf{Min: 0, Max: 0}
}

func buildSingleSize(param string) (WorSizeLimitConf, error) {
	var number, err = strconv.Atoi(param)
	if err != nil {
		return defaultConf(), err
	}
	return WorSizeLimitConf{Min: number, Max: number}, nil
}

func buildRange(param string) (WorSizeLimitConf, error) {
	var split = strings.Split(param, rangeSeparator)
	if len(split) != 2 {
		return defaultConf(), fmt.Errorf("word size limit value '%s' is invalid", param)
	}
	var min, errMin = strconv.Atoi(split[0])
	if errMin != nil {
		return defaultConf(), errMin
	}
	var max, errMax = strconv.Atoi(split[1])
	if errMax != nil {
		return defaultConf(), errMax
	}

	if min < max {
		return WorSizeLimitConf{Min: min, Max: max}, nil
	} else {
		return WorSizeLimitConf{Min: max, Max: min}, nil
	}
}
