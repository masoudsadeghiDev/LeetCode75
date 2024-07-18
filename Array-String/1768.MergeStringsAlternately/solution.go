package array_string

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {

	if word1 == "" && word2 == "" {
		return ""
	} else if word1 != "" && word2 == "" {
		return word1
	} else if word1 == "" && word2 != "" {
		return word2
	} else {
		p := 0           // pointer
		c1 := len(word1) // counter 1
		c2 := len(word2) // counter 2
		var sb strings.Builder

		for ; p < c1 && p < c2; p++ {
			sb.WriteString(fmt.Sprintf("%c", word1[p]))
			sb.WriteString(fmt.Sprintf("%c", (word2[p])))
		}

		if p >= c1 {
			for i := p; i < c2; i++ {
				sb.WriteString(fmt.Sprintf("%c", word2[i]))
			}
		} else if p >= c2 {
			for i := p; i < c1; i++ {
				sb.WriteString(fmt.Sprintf("%c", word1[i]))
			}
		}

		return sb.String()
	}
}
