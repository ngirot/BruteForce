// +build opencl

package conf

import (
	"gitlab.com/ngirot/blackcl"
)

func HasDeviceAvailable() bool {
	gpus, err := blackcl.GetDevices(blackcl.DeviceTypeGPU)
	if err != nil || len(gpus) == 0 {
		return false
	} else {
		return true
	}
}
