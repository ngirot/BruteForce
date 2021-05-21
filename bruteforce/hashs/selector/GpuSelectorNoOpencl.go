// +build !opencl

package selector

import (
	"errors"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
)

func BuildGpuHasherMap() (map[string]func() hashers.Hasher, error) {
	return nil, errors.New("GPU is not supported on this build or platform")
}
