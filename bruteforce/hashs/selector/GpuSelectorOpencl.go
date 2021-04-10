// +build opencl

package selector

import (
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
)

func BuildGpuHasherMap() (map[string]HasherInfos, error) {
	var hasherMap = make(map[string]HasherInfos)

	hasherMap["sha256"] = NewHasherInfos("SHA256", func() hashers.Hasher { return hashers.NewHasherGpuSha256() })

	return hasherMap, nil
}
