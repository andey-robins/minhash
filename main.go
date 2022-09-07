package main

import "flag"

func main() {
	var fPrime int
	flag.IntVar(&fPrime, "L", 2003, "the value of L in S = {0, 1, 2, 3, ... L-1}")
	flag.Parse()
}
