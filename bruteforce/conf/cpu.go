package conf

import "runtime"

func BestNumberOfGoRoutine() int {
	return runtime.NumCPU()*2 + 1;
}
