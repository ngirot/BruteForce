//go:build opencl
// +build opencl

package conf

import (
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers/gpu"
)

func HasDeviceAvailable() bool {
	_, err := gpu.GetDevice()
	if err != nil {
		return false
	} else {
		return true
	}
}
