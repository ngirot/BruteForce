// +build opencl

package selector

import (
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers/gpu"
)

func BuildGpuHasherMap() (map[string]func() hashers.Hasher, error) {
	var hasherMap = make(map[string]func() hashers.Hasher)

	hasherMap["sha256"] = func() hashers.Hasher { return gpu.NewHasherGpuSha256() }
	hasherMap["sha1"] = func() hashers.Hasher { return gpu.NewHasherGpuSha1() }

	return hasherMap, nil
}
