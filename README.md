# BruteForce [![Build Status](https://travis-ci.org/ngirot/BruteForce.svg?branch=master)](https://travis-ci.org/ngirot/BruteForce)
A simple brute forcer written in GO

## Usage

```
BruteForce --type [md5|sha256] --value <hash>
```

Example: 
```
./BruteForce --type sha256 --value ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad
Start brute-forcing...
Found : abc
In 0.012920 s
```

If you want to use a specific set of chars, write it on a simple txt file and use the --alphabet option
Example with a file alpha.data containing
```
abcdefghijklmnopqrstuvwxyz
```

You can launch it with:
```
./BruteForce --alphabet alpha.data --type sha256 --value 9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08
Start brute-forcing...
Found : test
In 0.081383 s
```

## Compilation
Just run in src folder:
```
go build -o BruteForce
```
