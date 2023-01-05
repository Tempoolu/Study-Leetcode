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

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 7, Left: nil, Right: nil}, Right: nil}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 5, Left: nil, Right: &TreeNode{Val: 6, Left: nil, Right: nil}},
	}
	fmt.Println(minDepth2(root))
}

// 思路和max depth一样，用全局变量作为答案
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1000000
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node.Left == nil && node.Right == nil {
			ans = int(math.Min(float64(depth), float64(ans)))
		}
		if node.Left != nil {
			dfs(node.Left, depth+1)
		}
		if node.Right != nil {
			dfs(node.Right, depth+1)
		}
	}
	dfs(root, 1)
	return ans
}
