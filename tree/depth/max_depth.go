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

// 二叉树最大深度
// 难度简单，但是我也没有自己做出来
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/solution/dong-hua-miao-shu-dfs-bfs-by-wfsbtong/

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 7, Left: nil, Right: nil}, Right: nil}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 5, Left: nil, Right: &TreeNode{Val: 6, Left: nil, Right: nil}},
	}
	fmt.Println(maxDepth2(root))
}

func maxDepth1(root *TreeNode) int {

	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			return 0
		}

		depth++
		lDepth := dfs(node.Left, depth)
		rDepth := dfs(node.Right, depth)
		return int(math.Max(float64(lDepth), float64(rDepth))) + 1
	}

	return dfs(root, 0)
}

// 这样思路很清晰：全局变量记录答案，每访问一层，都更新depth。然后比较每次的ans，选择最大那个。
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node.Left == nil && node.Right == nil {
			ans = int(math.Max(float64(ans), float64(depth)))
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
