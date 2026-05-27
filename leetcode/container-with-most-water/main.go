package main

import "fmt"

func maxArea(height []int) int {
	right := len(height) - 1
	left := 0
	area := 0

	for left < right {
		water := min(height[left], height[right]) * (right - left)

		area = max(area, water)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}

func main() {

	//driver
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
