package main

import "fmt"

func main() {
	canConstruct("aa", "ab")

	test := make(map[int]int)

	test[1]++
	fmt.Println(test[1])
	fmt.Println(test)
}

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)

	for _, v := range magazine {
		magazineMap[v]++
	}

	for _, val := range ransomNote {
		if magazineMap[val] == 0 {
			return false
		}

		magazineMap[val]--
	}

	return true
}
