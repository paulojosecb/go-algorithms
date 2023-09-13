package recursion

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	} else {
		return arr[0] + Sum(arr[1:])
	}
}

func Len(arr []int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return len(arr)
	} else {
		return 1 + Len(arr[1:])
	}
}
