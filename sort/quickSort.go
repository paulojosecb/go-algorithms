package sort

import "math/rand"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else if len(arr) == 2 {
		if arr[0] <= arr[1] {
			return arr
		} else {
			return append(arr[1:], arr[0])
		}
	} else {
		pivot := arr[rand.Intn(len(arr)-1)]
		lower := []int{}
		higher := []int{}

		for i := range arr {
			if arr[i] < pivot {
				lower = append(lower, arr[i])
			} else if arr[i] > pivot {
				higher = append(higher, arr[i])
			}
		}

		sortedLower := QuickSort(lower)
		sortedHigher := append([]int{pivot}, QuickSort(higher)...)

		return append(sortedLower, sortedHigher...)
	}

}
