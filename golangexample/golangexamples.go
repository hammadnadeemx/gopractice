package golangexamples

import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat string) string {
	var temp string
	var sc = "-"
	for i, _ := range sliceToConcat {
		temp += string(sliceToConcat[i]) + string(sc)
	}
	l := len(temp)
	return temp[:l-1]
}

func Encrypt(sliceToEncrypt string, ceaserCount int) string {
	carr := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var v int
	var encrypted string
	for _, chr := range sliceToEncrypt {
		for j, ele := range carr {
			if ele == string(chr) {
				v = (j + ceaserCount) % 26
				encrypted += carr[v]
				break
			}
		}
	}
	return encrypted
}
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
