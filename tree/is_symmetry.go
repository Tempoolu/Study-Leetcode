package main

import "fmt"

// 对称二叉树，简单，判断二叉树是否是对称的
// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: &TreeNode{Val: 4, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
	}
	fmt.Println(isSymmetry(root))
}

func isSymmetry(root *TreeNode) bool {

	var dfs func(lNode *TreeNode, rNode *TreeNode) bool
	dfs = func(lNode *TreeNode, rNode *TreeNode) bool {
		if lNode == nil && rNode == nil {
			return true
		}

		if lNode == nil || rNode == nil {
			return false
		}

		return lNode.Val == rNode.Val && dfs(lNode.Left, rNode.Right) && dfs(lNode.Right, rNode.Left)
	}

	return dfs(root.Left, root.Right)
}

func isSymmetry2(root *TreeNode) bool {
	ans := true
	var dfs func(lNode *TreeNode, rNode *TreeNode)
	dfs = func(lNode *TreeNode, rNode *TreeNode) {
		if lNode == nil && rNode == nil {
			return
		}
		if lNode == nil || rNode == nil {
			ans = false
			return
		}
		if lNode.Left != rNode.Right || lNode.Right != rNode.Left {
			ans = false
			return
		}
		dfs(lNode.Left, rNode.Right)
		dfs(lNode.Right, rNode.Left)
	}
	dfs(root.Left, root.Right)
	return ans
}
