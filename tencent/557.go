package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Let's take LeetCode contest"
	b := reverseWords(a)
	fmt.Println(b)
}

func reverseWords(s string) string {
	ans := strings.Split(s, " ")
	lengthA := len(ans)
	t := ""
	for i := 0; i < lengthA; i++ {
		index := []byte(ans[i])
		l := len(index)
		for i := 0; i < l/2; i++ {
			index[i], index[l-i-1] = index[l-i-1], index[i]
		}
		s = string(index)
		if i < lengthA-1 {
			s += " "
		}
		t += s
	}
	return t
}
