package main

import (
	"fmt"
	"maps"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	hash := map[string][]string{}

	for _, str := range strs {
		chars := []rune(str)
		slices.Sort(chars)
		key := string(chars)

		if _, exists := hash[key]; !exists {
			hash[key] = []string{}
		}

		hash[key] = append(hash[key], str)
	}

	return slices.Collect(maps.Values(hash))
}

func main() {
	//driver
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
