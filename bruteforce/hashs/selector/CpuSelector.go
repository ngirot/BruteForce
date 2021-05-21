package selector

import (
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
)

func BuildCpuHasherMap() map[string]func() hashers.Hasher {
	var hasherMap = make(map[string]func() hashers.Hasher)

	hasherMap["sha256"] = func() hashers.Hasher { return hashers.NewHasherSha256() }
	hasherMap["md5"] = func() hashers.Hasher { return hashers.NewHasherMd5() }
	hasherMap["sha1"] = func() hashers.Hasher { return hashers.NewHasherSha1() }
	hasherMap["sha512"] = func() hashers.Hasher { return hashers.NewHasherSha512() }
	hasherMap["bcrypt"] = func() hashers.Hasher { return hashers.NewHasherBcrypt() }

	return hasherMap
}
