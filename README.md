# Minhash
A version of the minhash algorithm for the Dynamic Big Data Course at Uwyo

## Dependencies

`go version go1.19 darwin/arm64`

## Running

`go run main.go --help`

### Arguments

`-L` -> Set the large prime number value
`-seed` -> Specify the seed for any random number generation. Specifying one will give deterministic output
`-class` -> Run the code with the example inputs and hashes given in class

## Building

`go build main.go`