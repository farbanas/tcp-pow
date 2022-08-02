package main

import (
	"crypto/sha256"
	"fmt"
	"tcp-pow/internal/calculators"
)

func main() {
	sc := calculators.NewSolutionCalculator(sha256.New())

	solution, err := sc.Calculate("abc", 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(solution)
}
