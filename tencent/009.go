package main

/**
妙
*/

func isPalindrome(x int) bool {
	y := x
	var num int
	for y != 0 {
		num = num*10 + y%10
		y = y / 10
	}
	if x < 0 || num != x {
		return false
	}
	return true
}
