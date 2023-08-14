package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(bubbleSort([]int{5, 3, 4, 1}))
	fmt.Println(selectionSort([]int{5, 3, 4, 1}))
}
