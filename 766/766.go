package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2}, {2, 2}}

	fmt.Println(isToeplitzMatrix(matrix))
}

func isToeplitzMatrix(matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n-1; i++ {
		for j := 0; j < m-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}

	return true
}

/**
	最优解
 */
func isToeplitzMatrix1(matrix [][]int) bool {
	if len(matrix) <= 1 || len(matrix[0]) <= 1 {
		return true
	}

	rows, column := len(matrix), len(matrix[0])
	// 横向检测
	for i := 0; i < column-1; i++ {
		x, y := 0, i
		for x+1 < rows && y+1 < column {
			if matrix[x][y] != matrix[x+1][y+1] {
				return false
			}
			x++
			y++
		}
	}
	// 竖向检测
	for i := 1; i < rows-1; i++ {
		x, y := i, 0
		for x+1 < rows && y+1 < column {
			if matrix[x][y] != matrix[x+1][y+1] {
				return false
			}
			x++
			y++
		}
	}
	return true
}