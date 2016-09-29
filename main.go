package main

import (
	"flag"
	"os"

	"github.com/co0p/matechallengefive/solutions"
)

func printUsage() {
	flag.PrintDefaults()
}

// main is responsible for delegating to the solutions based on the flags given
func main() {

	// handle flag parsing
	flag.Usage = printUsage
	numberPtr := flag.Int("solve", -1, "which sub challenge do you want me to solve (0-5)")
	flag.Parse()

	// are we in range?
	if *numberPtr < 0 || *numberPtr > 5 {
		printUsage()
		os.Exit(0)
	}

	// now go and solve it!
	switch *numberPtr {
	case 0:
		solutions.Zero()
	case 1:
		solutions.One()
	case 2:
		solutions.Two()
	case 3:
		solutions.Three()
	case 4:
		solutions.Four()
	case 5:
		solutions.Five()
	}
}
