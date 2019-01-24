package main

import "fmt"

func main() {
	i := []int{4, 1, 2, 1, 2}
	singleNumber(i)
}

func singleNumber(nums []int) int {
	out := 0

	for i := 0; i < len(nums); i++ {
		out = out ^ nums[i]
		fmt.Println(i, out)
	}
	return out
}
