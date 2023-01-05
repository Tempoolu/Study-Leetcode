package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将二叉树变成单向链表
// 简单，指针的概念很重要，需要再做一次
// https://leetcode.cn/problems/binode-lcci/

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: nil, Right: nil}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 5, Left: nil, Right: &TreeNode{Val: 6, Left: nil, Right: nil}},
	}
	// a := &TreeNode{Val: 1, Left: nil, Right: nil}
	fmt.Println(convertBiNode(root))
}

func convertBiNode(root *TreeNode) *TreeNode {

	dummy := &TreeNode{}
	pre := dummy // 记录前一个元素，也是学会使用指针

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			dfs(node.Left)
		}

		node.Left = nil
		pre.Right = node
		pre = node

		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)
	return dummy.Right //虚拟头的使用，学会使用指针
}
