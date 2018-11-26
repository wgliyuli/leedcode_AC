package main

import "fmt"

/**
	非负整数数组A，返回由A数组组成的前面为偶数后面为奇数的数组
 */
func main() {
	a := []int{3, 1, 2, 4}

	fmt.Println(sortArrayByParity(a))
}

func sortArrayByParity(A []int) []int {
	l := len(A)
	b := make([]int, l)
	j := 0
	k := l-1

	for i := 0; i < l; i++ {
		if A[i]%2 == 0 {
			b[j] = A[i]
			j++
		} else {
			b[k] = A[i]
			k--
		}
	}
	return b
}
