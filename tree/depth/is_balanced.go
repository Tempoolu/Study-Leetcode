package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 平衡二叉树，有关于深度，简单
// https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 7, Left: nil, Right: nil}, Right: nil}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 5, Left: nil, Right: &TreeNode{Val: 6, Left: nil, Right: nil}},
	}
	fmt.Println(isBalanced2(root))
}

// 思路：计算出最大深度和最小深度，看差距
func isBalanced2(root *TreeNode) bool {
	bigger := 0
	smaller := 10000000
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node.Left == nil && node.Right == nil {
			bigger = int(math.Max(float64(depth), float64(bigger)))
			smaller = int(math.Min(float64(depth), float64(bigger)))
		}
		if node.Left != nil {
			dfs(node.Left, depth+1)
		}
		if node.Right != nil {
			dfs(node.Right, depth+1)
		}
	}
	dfs(root, 1)
	fmt.Println(bigger, smaller)
	if bigger-smaller > 1 {
		return false
	}
	return true
}

func isBalanced1(root *TreeNode) bool {
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lDepth := dfs(node.Left)
		rDepth := dfs(node.Right)

		if lDepth != -1 && rDepth != -1 && abs(lDepth, rDepth) <= 1 {
			return int(math.Max(float64(lDepth), float64(rDepth))) + 1 // 使用math.Max(left_depth, right_depth)+1求最大深度
		}

		return -1
	}

	return dfs(root) != -1
}

func abs(n1 int, n2 int) int {
	if n1-n2 >= 0 {
		return n1 - n2
	} else {
		return n2 - n1
	}
}
