package main

import "fmt"

func validp(str string) bool {
	result := []string{}
	db := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, i := range str {
		if _, exists := db[string(i)]; exists {
			last := result[len(result)-1]
			result = result[:len(result)-1]
			if db[string(i)] != last {
				return false
			}
		} else {
			result = append(result, string(i))
		}
	}

	return len(result) == 0
}

func main() {

	fmt.Println(validp("()"))
	fmt.Println(validp("(]"))
}
