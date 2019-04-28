package algorithms

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isUnivalTree(root *TreeNode) bool {
	if root.Right != nil {
		if root.Val != root.Right.Val {
			return false
		}
		if !isUnivalTree(root.Right) {
			return false
		}
	}
	if root.Left != nil {
		if root.Val != root.Left.Val {
			return false
		}
		if !isUnivalTree(root.Left) {
			return false
		}
	}

	return true
}