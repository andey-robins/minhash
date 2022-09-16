# Minhash
A version of the minhash algorithm for the Dynamic Big Data Course at Uwyo

## Quick Reference

Easiest way to install go (that I've found): `brew install go`

Write up and discussion is in the file `writeup.pdf`

Run the example from class: `go run main.go -class`

Run a big example: `go run main.go -L 10007 -Q 1000 -seed 11`
    - Ommit the `-seed 11` for random results, keep and change the seed value for deterministic behavior

## Dependencies

`go version go1.19.1 darwin/arm64`

## Running

`go run main.go --help`

### Arguments

`-L` -> Set the large prime number value

`-Q` -> The number of hashes to generate

`-seed` -> Specify the seed for any random number generation. Specifying one will give deterministic output

`-class` -> Run the code with the example inputs and hashes given in class

`-help` -> Display help information

### Example

`go run main.go -L 10007 -Q 1000 -seed 11`

Provides an example output of:

```bash
Running program with values L=10007, Q=1000

Generated sparse columns
Constructed sparse matrix
Constructed hashes
Approximated Jaccard = 0.0049895092898855865
True Jaccard         = 0.0050380207787605
```

An approximation within 0.5% of the true Jaccard Similarity.

## RNG Information

All examples are generated using the golang rand package seeded either with Unix time or a specified seed value. Examples are computed on macOS Monterey 12.6, with arm architecture. Exact outputs may vary depending upon the system's handling of RNG.

## Building

`go build main.go`
