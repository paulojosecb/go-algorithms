package sort

func SelectSort(arr []int) []int {
	var sortedArr []int

	for 0 < len(arr) {
		lowerIndex := searchLowerIndex(arr)
		sortedArr = append(sortedArr, arr[lowerIndex])
		arr = append(arr[:lowerIndex], arr[lowerIndex+1:]...)
	}

	return sortedArr
}

func searchLowerIndex(arr []int) int {
	lowerIndex := 0

	for index := range arr {
		if arr[index] < arr[lowerIndex] {
			lowerIndex = index
		}
	}

	return lowerIndex
}
