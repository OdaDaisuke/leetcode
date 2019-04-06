package algorithms

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// var left *TreeNode
	// if root.Left != nil

	leftSide := maxDepth(root.Left) + 1
	rightSide := maxDepth(root.Right) + 1

	return max(leftSide, rightSide)
}

func max(left, right int) int {
	if left < right {
		return right
	}
	return left
}

// https://leetcode.com/submissions/detail/220435656/