package ch01_pure_fp

import (
	"fmt"
	. "bitbucket.org/lsheehan/fp-in-go/chapter1/01_fib"
)

func main() {
	fibSimple := FibSimple
	fmt.Printf("FibSimple: %v\n", fibSimple(8))

	fibMem := FibMemoized
	fmt.Printf("FibMemoized: %v\n", fibMem(8))


	fibChan := FibChanneled
	fmt.Printf("FibChanneled: %v\n", fibChan(8))

}


