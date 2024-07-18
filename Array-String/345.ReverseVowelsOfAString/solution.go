package array_string

import "fmt"

func main(){
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	p1 := 0 // Pointer that moves from the start of the string.
	p2 := len(s) - 1 // Pointer that moves from the end of the string.
	runes := []rune(s)
	for p1 < p2 {
		if isVowels(runes[p1]) {
			if isVowels(runes[p2]) { //Swap the runs if both of them are vowels.
				temp := runes[p1]
				runes[p1] = runes[p2]
				runes[p2] = temp
				p1++
			}
			p2--

		} else {
			p1++
		}
	}
	return string(runes)
}

func isVowels(r rune) bool {
	return r == 97 || // a
		r == 101 || // e
		r == 105 || // i
		r == 111 || // o
		r == 117 || // u
		r == 65 || // A
		r == 69 || // E
		r == 73 || // I
		r == 79 || // O
		r == 85 // U
}