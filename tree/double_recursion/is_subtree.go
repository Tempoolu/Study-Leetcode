package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 判断子树，简单，和same tree比较像，但是使用了双递归

func main() {
	// root1 := &TreeNode{
	// 	Val:   1,
	// 	Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: nil, Right: nil}, Right: &TreeNode{Val: 5, Left: nil, Right: nil}},
	// 	Right: &TreeNode{Val: 3, Left: nil, Right: nil},
	// }

	// root2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: nil, Right: nil}, Right: &TreeNode{Val: 5, Left: nil, Right: nil}}

	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: nil,
	}

	root2 := &TreeNode{Val: 1, Left: nil, Right: nil}
	fmt.Println(isSubtree(root1, root2))
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	finalAns := false // 双递归可以使用两个全局变量
	var dfs1 func(node *TreeNode)
	dfs1 = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs1(node.Left)
		dfs1(node.Right)
		if subRoot.Val == node.Val {
			ans := true // 第二个全局变量
			var dfs2 func(node, subNode *TreeNode)
			dfs2 = func(node, subNode *TreeNode) {
				if node == nil && subNode == nil {
					return
				}
				if node == nil || subNode == nil {
					ans = false
					return
				}

				if node.Val != subNode.Val {
					ans = false
					return
				}
				dfs2(node.Left, subNode.Left)
				dfs2(node.Right, subNode.Right)
			}
			dfs2(node, subRoot)
			if ans == true {
				finalAns = true
			}
		}
	}
	dfs1(root)
	return finalAns
}
