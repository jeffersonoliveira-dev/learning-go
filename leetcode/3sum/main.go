package main

import (
	"fmt"
	"slices"
)

func hasTripleSum(arr []int, target int) bool {
	n := len(arr)
	slices.Sort(arr)

	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1

		requiredSum := target - arr[i]

		for l < r {
			if arr[l]+arr[r] == requiredSum {
				return true
			}

			if arr[l]+arr[r] < requiredSum {
				l++
			}

			if arr[l]+arr[r] > requiredSum {
				r--
			}
		}
	}

	return false

}

func main() {

	arr := []int{1, 4, 45, 6, 10, 8}
	target := 13

	fmt.Println(hasTripleSum(arr, target))

}
