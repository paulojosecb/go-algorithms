package main

import (
	"fmt"

	"github.com/paulojosecb/go-algorithms/recursion"
)

func main() {
	testArr := []int{1, 2, 13, 4, 5}

	fmt.Print(recursion.Max(testArr))
}
