package main

import (
	"fmt"
)

func main() {
	A := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}

	fmt.Println(flipAndInvertImage(A))
}

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		k := len(A) - 1
		for j := 0; j < len(A); j++ {
			// 水平反转
			if j < len(A)/2 {
				A[i][j], A[i][k] = A[i][k], A[i][j]
				k--
			}

			// 反转图片
			if A[i][j] == 0 {
				A[i][j]++
			} else {
				A[i][j]--
			}
		}
	}

	return A
}
