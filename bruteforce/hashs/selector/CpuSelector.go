package selector

import (
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
)

func BuildCpuHasherMap() map[string]HasherInfos {
	var hasherMap = make(map[string]HasherInfos)

	hasherMap["sha256"] = NewHasherInfos("SHA256", func() hashers.Hasher { return hashers.NewHasherSha256() })
	hasherMap["md5"] = NewHasherInfos("MD5", func() hashers.Hasher { return hashers.NewHasherMd5() })
	hasherMap["sha1"] = NewHasherInfos("SHA1", func() hashers.Hasher { return hashers.NewHasherSha1() })
	hasherMap["sha512"] = NewHasherInfos("SHA512", func() hashers.Hasher { return hashers.NewHasherSha512() })
	hasherMap["bcrypt"] = NewHasherInfos("bcrypt (cost 10)", func() hashers.Hasher { return hashers.NewHasherBcrypt() })

	return hasherMap
}
