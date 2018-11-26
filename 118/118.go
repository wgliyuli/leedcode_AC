package main

import "fmt"

func main() {
	numRows := 5
	fmt.Println(generate(numRows))
}

func generate(numRows int) [][]int {
	A := make([][]int, 0)

	if numRows == 0 {
		return A
	}

	var B []int

	// 筑第一层
	B = make([]int, 1)
	B[0] = 1
	A = append(A, B)

	if numRows == 1 {
		return A
	}

	// 筑第二层
	B = make([]int, 2)
	B[0] = 1
	B[1] = 1
	A = append(A, B)
	if numRows == 2 {
		return A
	}

	// 筑后面n-2层
	for i := 3; i <= numRows; i++ {
		B = make([]int, i)
		B[0] = 1
		B[i-1] = 1
		t := 0
		for j := 1; j <= i-2; j++ {
			B[j] = A[i-2][t] + A[i-2][t+1]
			t++
		}
		A = append(A, B)
	}

	return A
}

/**
	官方解答
 */
func generate1(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	trigleArr := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		trigleArr[i] = make([]int, i+1)
	}
	trigleArr[0][0] = 1
	if numRows == 1 {
		return trigleArr
	}
	trigleArr[1][0] = 1
	trigleArr[1][1] = 1
	for i := 2; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				trigleArr[i][j] = 1
			} else {
				trigleArr[i][j] = trigleArr[i-1][j-1] + trigleArr[i-1][j]
			}
		}
	}
	return trigleArr

}
