package selector

import (
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers/cpu"
)

func BuildCpuHasherMap() map[string]func() hashers.Hasher {
	var hasherMap = make(map[string]func() hashers.Hasher)

	hasherMap["sha256"] = func() hashers.Hasher { return cpu.NewHasherSha256() }
	hasherMap["md5"] = func() hashers.Hasher { return cpu.NewHasherMd5() }
	hasherMap["sha1"] = func() hashers.Hasher { return cpu.NewHasherSha1() }
	hasherMap["sha512"] = func() hashers.Hasher { return cpu.NewHasherSha512() }
	hasherMap["bcrypt"] = func() hashers.Hasher { return cpu.NewHasherBcrypt() }

	return hasherMap
}
