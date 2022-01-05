package main

import "fmt"

func main() {

	BubbleSort([]int{3, 4, 5, 2, 1})

	inputArray := []int{6, 5, 3, 7, 2, 8, -1}
	minHeap := newMinHeap(inputArray)
	minHeap.sort(len(inputArray))
	minHeap.print()
	fmt.Scanln()
}
