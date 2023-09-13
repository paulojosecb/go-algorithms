package main

import (
	"fmt"

	"github.com/paulojosecb/go-algorithms/sort"
)

func main() {
	testArr := []int{5, 423, 31, 3, 1, 53}

	fmt.Print(sort.QuickSort(testArr))
}
