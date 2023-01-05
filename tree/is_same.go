package main

import "fmt"

// 判断是否是同一棵树，简单
// 完全是我自己做完的哦
// https://leetcode.cn/problems/same-tree/submissions/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 7, Left: nil, Right: nil}, Right: nil}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 5, Left: nil, Right: &TreeNode{Val: 6, Left: nil, Right: nil}},
	}

	root2 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 7, Left: nil, Right: nil}, Right: nil}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 5, Left: nil, Right: &TreeNode{Val: 6, Left: nil, Right: nil}},
	}
	fmt.Println(isSame(root1, root2))
}

func isSame(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if (root1 == nil || root2 == nil) || root1.Val != root2.Val {
		return false
	}

	return isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}

// 思路：用全局变量记录答案，如果存在val不相同，子节点存在不相同等问题，则改变ans为false。如果都不存在这些问题，则直接返回true
func isSame2(root1 *TreeNode, root2 *TreeNode) bool {
	ans := true
	var dfs func(root1 *TreeNode, root2 *TreeNode)
	dfs = func(root1 *TreeNode, root2 *TreeNode) {

		if root1.Val != root2.Val || (root1 != nil && root2 == nil) || (root1 == nil && root2 != nil) {
			ans = false
			return // 这里的return是结束一个递归，不会直接结束整个函数
		}
		if root1 == nil && root2 == nil {
			return
		}
		dfs(root1.Left, root2.Left)
		dfs(root1.Right, root2.Right)
	}
	dfs(root1, root2)
	return ans
}
