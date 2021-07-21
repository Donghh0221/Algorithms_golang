package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	alnumString := reg.ReplaceAllString(s, "")
	processedString := strings.ToLower(alnumString)

	length := len(processedString)
	n := length / 2
	for i := 0; i < n; i++ {
		if processedString[i] != processedString[(length-i-1)] {
			return false
		}
	}
	return true
}
