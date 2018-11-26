package main

import "fmt"

func main() {
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(transpose(A))
}

func transpose(A [][]int) [][]int {
	h, r := len(A), len(A[0])

	B := make([][]int, 0)

	for i := 0; i < r; i++ {
		C := make([]int, h)
		for j := 0; j < h; j++ {
			C[j] = A[j][i]
		}
		B = append(B, C)
	}

	return B
}

/**
	目前最优解
 */
func transpose1(A [][]int) [][]int {
	m, n := len(A), len(A[0])

	res := make([][]int, n)
	// 首先生成一个新的二维数组
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	// 依次插入数据
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = A[i][j]
		}
	}

	return res
}
