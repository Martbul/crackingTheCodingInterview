package main

import (
	//"fmt"
	"strconv"
	"strings"
)

//func main() {
//	res := compressedStringWIthBuilder("aabccccaaab")
//	fmt.Print(res)
//}

func stringCompression(s string) string {
	var compressedString string
	count := 1
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if s[i] == s[i+1] {
				count++
			} else {
				compressedString = compressedString + string(s[i]) + strconv.Itoa(count)
				count = 1
			}
		}
	}

	if len(s) < len(compressedString) {
		return s
	}
	return compressedString
}

func compressedStringWIthBuilder(s string) string {
	var builder strings.Builder
	count := 1
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if s[i] == s[i+1] {
				count++
			} else {
				builder.WriteByte(s[i])
				builder.WriteString(strconv.Itoa(count))
				count = 1
			}
		}
	}

	if builder.Len() >= len(s) {
		return s
	}
	return builder.String()
}
