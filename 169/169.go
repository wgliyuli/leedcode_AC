package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 3}

	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)

	res := nums[len(nums)/2]

	return res
}

/**
	最优解
 */
func majorityElement1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i, res := 1, nums[0]

	for _, v := range nums[1:] {
		// 如果数字相同就+1，如果不同就-1
		if v != res {
			i--
		} else {
			i++
		}

		if i <= 0 {
			res = v
			i = 1
		}
	}

	return res
}
