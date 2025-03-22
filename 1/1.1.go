package main

//
//import "fmt"

//func main() {
//	res := uniqueLetters("diod")
//	fmt.Println(res)
//}

func uniqueLetters(s string) bool {

	appearedLetters := make(map[string]int)

	for _, w := range s {
		_, ok := appearedLetters[string(w)]
		if ok {
			return false
		}

		appearedLetters[string(w)] = 1
	}
	return true
}
