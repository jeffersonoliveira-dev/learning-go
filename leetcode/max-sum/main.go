package main

import (
	"fmt"
)

func maxSum(arr []int, k int) int {
	n := len(arr)

	if n < k {
		fmt.Println("Invalid.")
		return -1
	}

	windowSum := 0

	for i := range k {
		windowSum += arr[i]
	}

	maxSum := windowSum

	for i := k; i < n; i++ {
		count := arr[i] - arr[i-k]
		windowSum += count
		maxSum = max(maxSum, windowSum)
	}

	return maxSum
}

func main() {

	arr := []int{5, 2, -1, 0, 3}

	k := 3

	// driver
	fmt.Println(maxSum(arr, k))

}
