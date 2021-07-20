package main

import (
	"regexp"
	"sort"
	"strings"
)

type letterLogFile struct {
	identifier string
	content    string
}

func (f *letterLogFile) tostring() string {
	return f.identifier + " " + f.content
}

type digitLogFile struct {
	identifier string
	content    string
}

func (f *digitLogFile) tostring() string {
	return f.identifier + " " + f.content
}

func reorderLogFiles(logs []string) []string {
	reg, _ := regexp.Compile("[0-9]+")

	var letters []letterLogFile
	var digits []digitLogFile

	for i := 0; i < (len(logs)); i++ {
		log := logs[i]
		identifier, content := strings.SplitN(log, " ", 2)[0], strings.SplitN(log, " ", 2)[1]
		if reg.Match([]byte(content)) {
			file := digitLogFile{identifier, content}
			digits = append(digits, file)
		} else {
			file := letterLogFile{identifier, content}
			letters = append(letters, file)
		}
	}

	sort.Slice(letters, func(i, j int) bool {
		return letters[i].identifier >= letters[j].identifier
	})
	sort.Slice(letters, func(i, j int) bool {
		return letters[i].content <= letters[j].content
	})

	var answer []string

	for i := 0; i < len(letters); i++ {
		file := letters[i]
		originalFile := file.tostring()
		answer = append(answer, originalFile)
	}
	for i := 0; i < len(digits); i++ {
		file := digits[i]
		originalFile := file.tostring()
		answer = append(answer, originalFile)
	}
	return answer
}
