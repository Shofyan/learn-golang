package main

func BubbleSort(arr []int) []int {

	if len(arr) < 1 {
		return []int{}
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}

		}

	}

	for _, v := range arr {
		println(v)
	}

	return arr
}
