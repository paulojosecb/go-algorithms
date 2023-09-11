package main

import (
	"fmt"

	"github.com/paulojosecb/go-algorithms/sort"
)

func main() {
	testArr := []int{4, 2, 1, 42, 5}

	fmt.Print(sort.SelectSort(testArr))
}
