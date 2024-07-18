package array_string

import (
	"regexp"
	"strings"
)


func reverseWords(s string) string {
	normalizedStr := NormalizeString(s)
	words := strings.Split(normalizedStr, " ")
	size := len(words)
	temp := ""
	for i := 0; i < size/2; i++ {
		temp = words[i]
		words[i] = words[size-i-1]
		words[size-i-1] = temp
	}
	return strings.Join(words, " ")
}

func NormalizeString(s string) string {
	trimmedStr := strings.TrimSpace(s)
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(trimmedStr, " ")
}
