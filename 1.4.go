package main

//
//import "fmt"

//func main() {
//	res := permPalindrome("aa")
//	fmt.Print(res)

//}

func permPalindrome(s string) bool {
	letters := make(map[string]int)
	lettersAppearedOddTimes := 0
	lettersAppeaderdEvenTimes := 0

	for i := 0; i < len(s); i++ {
		letter := string(s[i])
		if letter == " " {
			continue
		}
		val, ok := letters[letter]
		if !ok {
			letters[letter] = 1
			continue
		}
		newVal := val + 1
		letters[letter] = newVal

	}

	for _, appearances := range letters {
		if appearances%2 == 0 {
			lettersAppeaderdEvenTimes++
		} else {
			lettersAppearedOddTimes++
		}
	}
	if lettersAppearedOddTimes%2 != 0 || lettersAppeaderdEvenTimes%2 == 0 {
		return true
	} else {
		return false
	}

}
