package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int)
	for _, num := range nums {
		fmt.Println(hash[num])

		if hash[num] > 0 {
			return true
		} else {
			hash[num]++
		}
	}
	return false
}

func main() {

	nums := []int{1,2,3,}

	fmt.Println(containsDuplicate(nums))
}
