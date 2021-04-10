// +build !opencl

package selector

import (
	"errors"
)

func BuildGpuHasherMap() (map[string]HasherInfos, error) {
	return nil, errors.New("GPU is not supported on this build or platform")
}
