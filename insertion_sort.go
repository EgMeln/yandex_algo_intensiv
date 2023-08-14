package main

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		num := arr[i]
		j := i - 1
		for j >= 0 && num < arr[j] {
			arr[j+1], arr[j] = arr[j], arr[j+1]
			j--
		}
	}
	return arr
}
