package main

//import "fmt"

//func main() {
//	res := oneEditAway("pale", "bake")
//	fmt.Print((res))
//}

//INFO: THE BOOK SOLUTION

func oneEditAway(s1 string, s2 string) bool {
	if len(s1) == len(s2) {
		return oneEditReplace(s1, s2)
	} else if len(s1)+1 == len(s2) {
		return oneEditInsert(s1, s2)
	} else if len(s1)-1 == len(s2) {
		return oneEditInsert(s2, s1)
	}
	return false
}

func oneEditReplace(s1 string, s2 string) bool {
	foundDifference := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if foundDifference {
				return false
			}

			foundDifference = true
		}
	}
	return true
}

func oneEditInsert(s1 string, s2 string) bool {
	var index1 int
	var index2 int
	for index2 < len(s2) && index1 < len(s1) {
		if s1[index1] != s2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}
	return true
}

// INFO: This is my failed try
func oneAway(s1 string, s2 string) bool {
	wordTable1 := make(map[string]int)
	wordTable2 := make(map[string]int)
	letterSum1 := 0
	letterSum2 := 0

	for i := 0; i < len(s1); i++ {
		letter := string(s1[i])
		val, ok := wordTable1[letter]
		if !ok {
			wordTable1[letter] = 1
		}

		wordTable1[letter] = val + 1
	}

	for i := 0; i < len(s2); i++ {
		letter := string(s2[i])
		val, ok := wordTable2[letter]
		if !ok {
			wordTable2[letter] = 1
		}

		wordTable2[letter] = val + 1

	}

	for _, num := range wordTable1 {
		letterSum1 += num

	}

	for _, num := range wordTable2 {
		letterSum2 += num

	}

	//fmt.Print(wordTable1)
	//fmt.Print(wordTable2)

	//WARN: FUCK IT UP, JUST SUMMING LETTERS NOTHING MORE
	if letterSum1 > letterSum2 {
		if letterSum1-letterSum2 > 1 {
			return false
		} else {
			return true
		}
	}

	if letterSum1 < letterSum2 {
		if letterSum2-letterSum1 > 1 {
			return false
		} else {
			return true
		}
	}
	return true
}
