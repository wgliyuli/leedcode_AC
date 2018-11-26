package main

import "fmt"

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	r := 1
	c := 4

	fmt.Println(matrixReshape(nums, r, c))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	n, m := len(nums), len(nums[0])
	if n*m != r*c {
		return nums
	}
	A := make([][]int, r)

	// 生成矩阵
	for i := 0; i < r; i++ {
		A[i] = make([]int, c)
	}

	// 取出数据
	B := make([]int, n*m)
	t := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			B[t] = nums[i][j]
			t++
		}
	}

	// 插入数据
	t = 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			A[i][j] = B[t]
			t++
		}
	}
	return A
}

/**
	最优解
 */
func matrixReshape1(nums [][]int, r int, c int) [][]int {
	n, m:=len(nums), len(nums[0])
	if n*m!=r*c{
		return nums
	}
	res:=make([][]int, r)
	for i:=0;i<r;i++{
		res[i] = make([]int,c)
	}

	// 精髓
	q,a:=0,0
	for _,v:=range nums{
		for _,v2:=range v{
			res[q][a] = v2
			a++
			if a==c{
				a=0
				q++
			}
		}
	}
	return res
}