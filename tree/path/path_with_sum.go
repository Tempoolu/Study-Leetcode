package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 路径求和等于某一个值，中等，自己做出来
// https://leetcode.cn/problems/paths-with-sum-lcci/

func main() {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil},
	}

	fmt.Println(pathSum(root, 3))
}

func pathSum(root *TreeNode, sum int) int {
	count := 0

	var dfs1 func(node *TreeNode)
	dfs1 = func(node *TreeNode) {
		if node == nil {
			return
		}
		var dfs2 func(node *TreeNode, path []int)
		dfs2 = func(node *TreeNode, path []int) {

			if node == nil {
				return
			}

			dfs2(node.Left, append(path, node.Val))
			dfs2(node.Right, append(path, node.Val))

			if addPath(path)+node.Val == sum {
				count++
				return
			}
			if addPath(path)+node.Val > sum {
				return
			}
			if node.Left == nil && node.Right == nil {
				return
			}
		}

		dfs1(node.Left)
		dfs1(node.Right)
		dfs2(node, []int{})
	}
	dfs1(root)
	return count
}

func addPath(path []int) int {
	res := 0
	for _, v := range path {
		res += v
	}
	return res
}
