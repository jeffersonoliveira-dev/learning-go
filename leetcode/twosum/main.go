package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// create hash
	// set loop
	// check if the rest of the count is on the hash
	// if the rest of the count is not on the hash, set the new value with the index
	hash := make(map[int]int)
	for index, value := range nums {
		rest := target - value

		if _, exists := hash[rest]; exists {
			ind := hash[rest]
			return []int{ind, index}
		} else {
			hash[value] = index
		}
	}

	return []int{}
}

func main() {

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
