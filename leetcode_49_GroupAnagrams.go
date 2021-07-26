package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		sorted_str := strings.Join(s, "")
		m[sorted_str] = append(m[sorted_str], str)
	}
	answer := make([][]string, 0, 0)
	for _, v := range m {
		answer = append(answer, v)
	}
	return answer
}
