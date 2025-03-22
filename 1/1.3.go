package main

import "fmt"

//func main() {
//	res := URLify("MR sd e % sgsf sdfd sdof w89yr00")
//	//res := URLify("MR sd ef  % sgsfg    sdfd  sdof w89yr00")
//	fmt.Print(res)

//}

func URLify(s string) string {
	var res string

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			//res = res + "%20"

			res = fmt.Sprintf("%s%s", res, "%20")
			continue
		}

		//res = res + string(s[i])
		res = fmt.Sprintf("%s%s", res, string(s[i]))
	}
	return res

}
