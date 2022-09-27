//go:build opencl
// +build opencl

package gpu

import "gitlab.com/ngirot/blackcl"

func GetDevice() (*blackcl.Device, error) {
	gpus, err := blackcl.GetDevices(blackcl.DeviceTypeGPU)
	if len(gpus) != 0 && err == nil {
		return gpus[0], nil
	}

	device, err := blackcl.GetDevice(blackcl.DeviceTypeGPU)
	return device, err
}
