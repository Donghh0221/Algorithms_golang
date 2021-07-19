package main

func reverseString(s []byte) {
	length := len(s)
	p := length / 2
	for i := 0; i < p; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
	return
}
