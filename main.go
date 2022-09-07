package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information.")
	}

	var fPrime int
	var help, class bool
	flag.IntVar(&fPrime, "L", 11, "the value of L in S = {0, 1, 2, 3, ... L-1}")
	flag.BoolVar(&help, "help", false, "print help information")
	flag.BoolVar(&class, "class", false, "set values to default for class example")
	flag.Parse()

	if help {
		pad := func() {
			fmt.Printf("\n\n")
		}
		pad()
		fmt.Println(" Welcome to the minhash explorer tool!")
		fmt.Print(" This code is licensed under GPLv3")
		pad()
		fmt.Println(" Args:")
		fmt.Println("  -L       The upper bound in defining S = {0, 1, 2, ..., L-1}")
		fmt.Println("  -class   Automatically set defaults to the class example")
		fmt.Println("  -help    Display this help text :)")
		pad()
	}
}
