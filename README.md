# BruteForce [![Build Status](https://travis-ci.org/ngirot/BruteForce.svg?branch=master)](https://travis-ci.org/ngirot/BruteForce)
A simple brute force software written in GO

## Usage

```
BruteForce --type [md5|sha256|sha512|sha1] --value <hash>
```

Example: 
```
./BruteForce --type sha256 --value ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad
Start brute-forcing (sha256)...
Found: abc in 0 s
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
If you want to use a salt you can simply add it with the salt parameter
Example:
```
./BruteForce --value cb2537e62f7f7358cb5b6a812c7d3c7d95d87733 --type sha1 --salt 1234567890
Start brute-forcing (sha1)...
Found: zzzz in 1 s


```