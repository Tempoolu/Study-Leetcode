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

// 二叉搜索树，简单，特征是从左到右越来越大，如果进行中序遍历，则是有序的
// https://leetcode.cn/problems/minimum-distance-between-bst-nodes/

func main() {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil},
	}

	fmt.Println(minDiffInBST(root))
}

// 使用两个全局变量
func minDiffInBST(root *TreeNode) int {
	min := 100000
	pre := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		min = int(math.Min(float64(node.Val-pre), float64(min)))
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)

	return min
}

func minDiffInBST2(root *TreeNode) int {
	res := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	return findMin(res)
}

func findMin(arr []int) int {
	min := 100000
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < min {
			min = arr[i+1] - arr[i]
		}
	}
	return min
}
