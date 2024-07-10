package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abc"
	str2 := "pqr"
	output := gcdOfStrings(str1, str2)
	fmt.Printf("Input %v, %v and output is %v", str1, str2, output)
}

func gcdOfStrings(str1 string, str2 string) string {

	l1 := len(str1)
	l2 := len(str2)
	var sb strings.Builder
	tptr := l2

	for i := 0; i < l2; i++ {
		sub := str2[0:tptr]
		sl := len(sub)

		if l1%sl == 0 && l2%sl == 0 {
			d2 := l2 / sl
			sb.Reset()
			for j := 0; j < d2; j++ {
				sb.WriteString(sub)
			}
			if sb.String() == str2 {
				sb.Reset()
				d1 := l1 / sl
				for j := 0; j < d1; j++ {
					sb.WriteString(sub)
				}
				if sb.String() == str1 {
					return sub
				}
			}
		}

		tptr--
	}
	return ""
}
