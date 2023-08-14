package main

func findMaxPodsum(arr []int) int {
	maxSum := 0
	currVal := 0
	for i := 0; i < len(arr); i++ {
		if currVal+arr[i] < 0 {
			currVal = 0
		} else {
			currVal += arr[i]
			if currVal > maxSum {
				maxSum = currVal
			}
		}
	}
	return maxSum
}
