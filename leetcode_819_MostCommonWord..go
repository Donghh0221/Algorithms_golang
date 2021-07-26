package main

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	lower := strings.ToLower(paragraph)
	re, _ := regexp.Compile(`[^\w]`)
	processed := re.ReplaceAllString(lower, " ")
	words := strings.Fields(processed)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	for _, banned_word := range banned {
		delete(m, banned_word)
	}
	answer := ""
	count := 0
	for k, v := range m {
		if v > count {
			count = v
			answer = k
		}
	}
	return answer
}
