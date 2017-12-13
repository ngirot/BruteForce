# BruteForce [![Build Status](https://travis-ci.org/ngirot/BruteForce.svg?branch=master)](https://travis-ci.org/ngirot/BruteForce)
A simple brute force software written in GO

## Usage

```
BruteForce --type [md5|sha256|sha512|sha1] --value <hash>
```

Example: 
```
./BruteForce --type sha256 --value ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad
Start brute-forcing...
Found: abc
In 0.012920 s
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
Start brute-forcing...
Found: test
In 0.081383 s
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
Start brute-forcing...
Found: test
In 0.000531 s
```