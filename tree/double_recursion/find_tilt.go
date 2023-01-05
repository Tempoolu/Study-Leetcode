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

// 找到坡度，简单，主要写了使用全局变量和不使用全局变量，以及双递归
// 双递归主要用于，在中间一个节点要求什么东西
// https://leetcode.cn/problems/binary-tree-tilt/submissions/

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
		Right: &TreeNode{Val: 2, Left: nil, Right: nil},
	}

	fmt.Println(findTilt2(root))
}

// 不用全局变量
func findTilt(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 以下
		return int(math.Abs(float64(add(node.Left))-float64(add(node.Right)))) + findTilt(node.Left) + findTilt(node.Right)
	}
	return dfs(root)
}

// 用全局变量
func findTilt2(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 以下
		dfs(node.Left)
		dfs(node.Right)
		ans += int(math.Abs(float64(add(node.Left)) - float64(add(node.Right))))
	}
	dfs(root)
	return ans
}

// 不用全局变量，返回值就是最终值，在return要想清楚
func add(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return node.Val + dfs(node.Left) + dfs(node.Right)
	}
	return dfs(root)
}

// 用全局变量，不需要返回值
func add2(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res += node.Val
	}
	dfs(root)
	return res
}
