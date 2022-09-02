# BruteForce [![Build Status](https://travis-ci.org/ngirot/BruteForce.svg?branch=master)](https://travis-ci.org/ngirot/BruteForce)
A simple brute force software written in GO

## Usage

```
BruteForce --type [md5|sha256|sha512|sha1|bcrypt|ripemd160] --value <hash>
```

Example: 
```
./BruteForce --type sha256 --value 88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589
Start brute-forcing (sha256)...
Found: abc in 1 s
```

### How to use a specific char set 
If you want to use a specific set of chars, write it on a simple txt file and use the --alphabet option
Example with a file alpha.data containing:
```
abcdefghijklmnopqrstuvwxyz
```

You can launch it with:
```
./BruteForce --alphabet alpha.data --type sha256 --value 9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08
Start brute-forcing (sha256)...
Found: test in 0 s
```

### How to use a dictionary file
If you want to use a file containing all words to be tested, write them on a simple file (one per line) and use the --dictionary option.
Example with a file dic.data containing:

```
dragon
butterfly
test
```

You can launch it with:
```
./BruteForce --dictionary dic.data --type sha256 --value 9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08
Start brute-forcing (sha256)...
Found: test in 0 s
```

### How to use a salt
If you want to use a salt you can simply add it with the salt parameter (before and/or after)
Example:
```
./BruteForce --value cb2537e62f7f7358cb5b6a812c7d3c7d95d87733 --type sha1 --salt-after 1234567890
Start brute-forcing (sha1)...
Found: zzzz in 3 s
```
```
./BruteForce --value ef95124ec674cea240e9dd02c86ba3670e2ee5a2 --type sha1 --salt-before 1234567890
Start brute-forcing (sha1)...
Found: zzzz in 3 s
```

### Gpu support (beta)
#### Use GPU support
Currently, only SHA-256 is supported with this option, and only for Windows and Linux builds, with an OpenCL compatible device and driver.

And if you want to use GPU to compute hashs, just use the GPU parameter:
```
./BruteForce --type sha256 --value 9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08 --gpu
Start brute-forcing '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08' (sha256)...
Found: test in 9 s
```

#### Build with GPU support
If you want to build with GPU support, you must use the command `go build -tags opencl`.
You will also need header and lib files available.
- Linux : You just have to install the right package on your Linux system (example `dnf install opencl-headers` on Fedora)
- Windows : download the OpenCL package (https://github.com/GPUOpen-LibrariesAndSDKs/OCL-SDK/releases) and use GO flags
  `CGO_CPPFLAGS=-I c:\OCL_SDK_Light\include` and `CGO_LDFLAGS=-L c:\OCL_SDK_Light\lib\x86_64`


