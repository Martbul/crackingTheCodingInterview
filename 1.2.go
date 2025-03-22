package main

import (
	//	"fmt"
	"sort"
)

//func main() {
//	res := isPErmutationOfString("dog", "god ")
//	fmt.Println(res)
//}

func isPErmutationOfString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1 = sortString(s1)
	s2 = sortString(s2)

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true

}

func sortString(input string) string {
	runes := []rune(input)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
