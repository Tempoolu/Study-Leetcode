package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树剪去枝，中等，自己没做出来
// https://leetcode.cn/problems/binary-tree-pruning/
// 属于后序遍历的经典题目：为什么使用后序遍历？如果需要使用节点本身的值或参数（这里是指针）来传递给下一层，则使用后序遍历
// 这里是将父节点指向空，只能在父节点处操作，不应遍历到下一层。因此使用后序遍历

// 前序遍历常常和参数扩展结合，例如参数扩展路径、参数拓展深度等
func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: nil},
		Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: nil, Right: nil}, Right: nil},
	}

	fmt.Println(pruneTree(root))

}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		if node.Val == 0 && node.Left == nil && node.Right == nil {
			return nil
		} else {
			return node
		}
	}
	dfs(root)
	return root
}
