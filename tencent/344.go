package main

import "fmt"

func main() {
	a := "hello"
	b := []byte(a)
	reverseString(b)
	t := string(b[:])
	fmt.Println(t)
}

func reverseString(s []byte) {
	t := len(s)
	for i := 0; i < t/2; i++ {
		s[i], s[t-i-1] = s[t-i-1], s[i]
	}
}
