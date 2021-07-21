package main

import (
	"strconv"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	rawString := strconv.Quote(s)
	var que []rune
	answer := 0

	if len(rawString) == 0 {
		return answer
	}
	for _, cha := range s {
		if strings.ContainsRune(string(que), cha) {
			idx := strings.IndexRune(string(que), cha)
			que = que[idx+1:]
			que = append(que, cha)
		} else {
			que = append(que, cha)
		}
		if answer < len(que) {
			answer = len(que)
		}
	}
	return answer
}
