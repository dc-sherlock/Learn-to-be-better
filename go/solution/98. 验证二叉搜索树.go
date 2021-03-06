/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var last int

func isValidBinarySearchTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isValidBinarySearchTree(root.Right) {
		if last > root.Val {
			last = root.Val
			return isValidBinarySearchTree(root.Left)
		}
	}
	return false
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	last = int(^uint(0) >> 1)
	return isValidBinarySearchTree(root)
}
