package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/andey-robins/minhash/class"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Run with -help for help information.")
	}

	var prime, seed int
	var help, classOnly bool
	flag.IntVar(&prime, "L", 11, "the value of L in S = {0, 1, 2, 3, ... L-1}")
	flag.IntVar(&seed, "seed", int(time.Now().UnixNano()), "the seed for random number generation")
	flag.BoolVar(&help, "help", false, "print help information")
	flag.BoolVar(&classOnly, "class", false, "set values to default for class example")
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
		fmt.Println("  -L        The upper bound in defining S = {0, 1, 2, ..., L-1}")
		fmt.Println("  -seed     The seed for all RNG used in this code. Defaults to Unix Time if none provided")
		fmt.Println("  -class    Automatically set defaults to the class example")
		fmt.Println("  -help     Display this help text :)")
		pad()
		return
	}

	if classOnly {
		class.ClassDriver()
		return
	}

	driver(prime, seed)
}

func driver(l, seed int) {
	fmt.Printf("Running program with values L=%v, seed=%v\n", l, seed)
}
