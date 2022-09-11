# Minhash
A version of the minhash algorithm for the Dynamic Big Data Course at Uwyo

## Dependencies

`go version go1.19.1 darwin/arm64`

## Running

`go run main.go --help`

### Arguments

`-L` -> Set the large prime number value
`-Q` -> The number of hashes to generate
`-seed` -> Specify the seed for any random number generation. Specifying one will give deterministic output
`-class` -> Run the code with the example inputs and hashes given in class

### Example

`go run main.go -L 10007 -Q 1000 -seed 11`

Provides an example output of:

```bash
Running program with values L=10007, Q=1000

Generated sparse columns
Constructed sparse matrix
Constructed hashes
Approximated Jaccard = 0.010362694300518135
True Jaccard         = 0.010309278350515464
```

An approximation within 0.5% of the true Jaccard Similarity.

## RNG Information

All examples are generated using the golang rand package seeded either with Unix time or a specified seed value. Examples are computed on macOS Monterey 12.5, with arm architecture. Exact outputs may vary depending upon the system's handling of RNG.

## Building

`go build main.go`
