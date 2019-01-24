package main

import "math"

/**
递归
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		left_height := maxDepth(root.Left)
		right_height := maxDepth(root.Right)
		return int(math.Max(float64(left_height), float64(right_height))) + 1
	}
}
