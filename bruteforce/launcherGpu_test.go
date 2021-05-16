// +build opencl

package bruteforce

import (
	"github.com/ngirot/BruteForce/bruteforce/conf"
	"testing"
)

var gpuConfiguration = conf.NewProcessingUnitConfiguration(true)

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabet_GPU_IntegrationTest(t *testing.T) {
	runAlphabetTest(t, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabetWithSaltAfter_GPU_IntegrationTest(t *testing.T) {
	runAlphabetSaltAfter(t, "d43f1b69d87466016e70505bf91c0fd2f075a905b6d18b3bbe2f60ae2ca3dac6", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromAlphabetWithSaltBefore_GPU_IntegrationTest(t *testing.T) {
	runAlphabetSaltBefore(t, "73963907b19c38e6cf8d4c9e4c08d966269abd5a17dca0171804be9121bd525c", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionary_GPU_IntegrationTest(t *testing.T) {
	runDictionary(t, "3eb95849b1228d1f28f57e60aa691e78f295245e79681e3bbb7a5807d4e01ed6", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionaryWithSaltAfter_GPU_IntegrationTest(t *testing.T) {
	runDictionartSaltAfter(t, "a6c624948b941bd972cf11d58af3cb5bb3ed326eb8107fe7bb5cb4f549706455", "sha256", cpuConfiguration)
}

func Test_Launcher_ShouldFindSimpleSha256HashFromDictionaryWithSaltBefore_GPU_IntegrationTest(t *testing.T) {
	runDictionarySaltBefore(t, "75e1aea225bdb80135b82dce07dfb302a2654bd680113eaf05807017e6a83fe4", "sha256", cpuConfiguration)
}
