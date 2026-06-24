package main

import (
	"slices"
	"fmt"
)

func isAnagram(first string, second string) bool {
	f := []byte(first)
	s := []byte(second)

	slices.Sort(f)
	slices.Sort(s)

	finalFirst := string(f)
	finalSecond := string(s)



	return finalFirst == finalSecond
}

func main(){

	first := "anagram"
	second := "nagaram"

	fmt.Println(isAnagram(first, second))
}
