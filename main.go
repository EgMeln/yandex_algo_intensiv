package main

import "fmt"

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
	graph := &node{Val: 2, R: &node{Val: 3}}

	sum := 0
	dfs(graph, &sum)
	fmt.Println(sum)
}
