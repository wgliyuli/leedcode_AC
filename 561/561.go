package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(A))
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}

	return sum
}
