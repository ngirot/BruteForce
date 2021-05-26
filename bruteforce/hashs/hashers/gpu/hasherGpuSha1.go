// +build opencl

package gpu

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/ngirot/BruteForce/bruteforce/hashs/hashers"
	"gitlab.com/ngirot/blackcl"
	"strings"
)

type hasherGpuSha1 struct {
	device           *blackcl.Device
	kernelDictionary *blackcl.Kernel
	kernelAlphabet   *blackcl.Kernel
	endianness       binary.ByteOrder
}

func NewHasherGpuSha1() hashers.Hasher {
	gpus, err := blackcl.GetDevices(blackcl.DeviceTypeGPU)
	if err == nil {
		for _, device := range gpus {
			device.AddProgram(buildKernelsSha1())
			var endianness = detectEndianness(device, "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3")
			return &hasherGpuSha1{device, device.Kernel(genericKernelCryptName), device.Kernel(genericKernelCryptAndWorderName), endianness}
		}
	}

	return nil
}

func (h *hasherGpuSha1) Example() string {
	return hex.EncodeToString(h.Hash([]string{"1234567890"})[0])
}

func (h *hasherGpuSha1) DecodeInput(data string) []byte {
	var result, _ = hex.DecodeString(data)
	return result
}

func (h *hasherGpuSha1) Hash(datas []string) [][]byte {
	return genericHashWithGpu(h.device, h.kernelDictionary, h.endianness, datas, 20)
}

func (h *hasherGpuSha1) IsValid(data string) bool {
	return hashers.GenericBase64Validator(h, data)
}

func (h *hasherGpuSha1) Compare(transformedData []byte, referenceData []byte) bool {
	return bytes.Equal(transformedData, referenceData)
}

func (h *hasherGpuSha1) ProcessWithWildcard(charSet []string, saltBefore string, saltAfter string, numberOfWildCards int, expectedDigest string) string {
	return genericProcessWithGpu(h.device, h.kernelAlphabet, h.endianness,
		charSet, saltBefore, saltAfter, numberOfWildCards, expectedDigest)
}

// https://github.com/Fruneng/opencl_sha_al_im
func buildKernelsSha1() string {
	var parametrized = `
#ifdef cl_khr_byte_addressable_store
#pragma OPENCL EXTENSION cl_khr_byte_addressable_store : disable
#endif

#ifdef cl_nv_pragma_unroll
#define NVIDIA
#pragma OPENCL EXTENSION cl_nv_pragma_unroll : enable
#endif

#ifdef NVIDIA
inline uint SWAP32(uint x)
{
	x = rotate(x, 16U);
	return ((x & 0x00FF00FF) << 8) + ((x >> 8) & 0x00FF00FF);
}
#else
#define SWAP32(a)	(as_uint(as_uchar4(a).wzyx))
#endif

#define K0  0x5A827999
#define K1  0x6ED9EBA1
#define K2  0x8F1BBCDC
#define K3  0xCA62C1D6

#define H1 0x67452301
#define H2 0xEFCDAB89
#define H3 0x98BADCFE
#define H4 0x10325476
#define H5 0xC3D2E1F0

#ifndef uint32_t
#define uint32_t unsigned int
#endif

uint32_t SHA1CircularShift(int bits, uint32_t word)
{
	return ((word << bits) & 0xFFFFFFFF) | (word) >> (32 - (bits));
}

void hash(global char *plain_key, uint *digest, uint ulen) {
int t, gid, msg_pad;
    int stop, mmod;
    uint i, item, total;
    uint W[80], temp, A,B,C,D,E;
	int current_pad;

	msg_pad=0;

	total = ulen%64>=56?2:1 + ulen/64;

	//printf("ulen: %u total:%u\n", ulen, total);

    digest[0] = 0x67452301;
	digest[1] = 0xEFCDAB89;
	digest[2] = 0x98BADCFE;
	digest[3] = 0x10325476;
	digest[4] = 0xC3D2E1F0;
	for(item=0; item<total; item++)
	{

		A = digest[0];
		B = digest[1];
		C = digest[2];
		D = digest[3];
		E = digest[4];

	#pragma unroll
		for (t = 0; t < 80; t++){
		W[t] = 0x00000000;
		}
		msg_pad=item*64;
		if(ulen > msg_pad)
		{
			current_pad = (ulen-msg_pad)>64?64:(ulen-msg_pad);
		}
		else
		{
			current_pad =-1;		
		}

		//printf("current_pad: %d\n",current_pad);
		if(current_pad>0)
		{
			i=current_pad;

			stop =  i/4;
			//printf("i:%d, stop: %d msg_pad:%d\n",i,stop, msg_pad);
			for (t = 0 ; t < stop ; t++){
				W[t] = ((uchar)  plain_key[msg_pad + t * 4]) << 24;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 1]) << 16;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 2]) << 8;
				W[t] |= (uchar)  plain_key[msg_pad + t * 4 + 3];
				//printf("W[%u]: %u\n",t,W[t]);
			}
			mmod = i % 4;
			if ( mmod == 3){
				W[t] = ((uchar)  plain_key[msg_pad + t * 4]) << 24;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 1]) << 16;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 2]) << 8;
				W[t] |=  ((uchar) 0x80) ;
			} else if (mmod == 2) {
				W[t] = ((uchar)  plain_key[msg_pad + t * 4]) << 24;
				W[t] |= ((uchar) plain_key[msg_pad + t * 4 + 1]) << 16;
				W[t] |=  0x8000 ;
			} else if (mmod == 1) {
				W[t] = ((uchar)  plain_key[msg_pad + t * 4]) << 24;
				W[t] |=  0x800000 ;
			} else /*if (mmod == 0)*/ {
				W[t] =  0x80000000 ;
			}
			
			if (current_pad<56)
			{
				W[15] =  ulen*8 ;
				//printf("w[15] :%u\n", W[15]);
			}
		}
		else if(current_pad <0)
		{
			if( ulen%64==0)
				W[0]=0x80000000;
			W[15]=ulen*8;
			//printf("w[15] :%u\n", W[15]);
		}

		

		for (t = 16; t < 80; t++)
		{
			W[t] = SHA1CircularShift(1, W[t - 3] ^ W[t - 8] ^ W[t - 14] ^ W[t - 16]);
		}

		for (t = 0; t < 20; t++)
		{
			temp = SHA1CircularShift(5, A) +
				((B & C) | ((~B) & D)) + E + W[t] + K0;
			temp &= 0xFFFFFFFF;
			E = D;
			D = C;
			C = SHA1CircularShift(30, B);
			B = A;
			A = temp;
		}

		for (t = 20; t < 40; t++)
		{
			temp = SHA1CircularShift(5, A) + (B ^ C ^ D) + E + W[t] + K1;
			temp &= 0xFFFFFFFF;
			E = D;
			D = C;
			C = SHA1CircularShift(30, B);
			B = A;
			A = temp;
		}

		for (t = 40; t < 60; t++)
		{
			temp = SHA1CircularShift(5, A) +
				((B & C) | (B & D) | (C & D)) + E + W[t] + K2;
			temp &= 0xFFFFFFFF;
			E = D;
			D = C;
			C = SHA1CircularShift(30, B);
			B = A;
			A = temp;
		}

		for (t = 60; t < 80; t++)
		{
			temp = SHA1CircularShift(5, A) + (B ^ C ^ D) + E + W[t] + K3;
			temp &= 0xFFFFFFFF;
			E = D;
			D = C;
			C = SHA1CircularShift(30, B);
			B = A;
			A = temp;
		}

		digest[0] = (digest[0] + A) & 0xFFFFFFFF;
		digest[1] = (digest[1] + B) & 0xFFFFFFFF;
		digest[2] = (digest[2] + C) & 0xFFFFFFFF;
		digest[3] = (digest[3] + D) & 0xFFFFFFFF;
		digest[4] = (digest[4] + E) & 0xFFFFFFFF;

		//for(i=0;i<80;i++)
			//printf("W[%u]: %u\n", i,W[i] );

		//printf("%u\n",  digest[0]);
		//printf("%u\n",  digest[1]);
		//printf("%u\n",  digest[2]);
		//printf("%u\n",  digest[3]);
		//printf("%u\n",  digest[4]);
	}
}

___KERNELS___

}
`
	return strings.ReplaceAll(parametrized, "___KERNELS___", buildGenericKernel(20))
}
