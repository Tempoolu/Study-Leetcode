package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将树镜像，简单，自己做出来的
// https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/submissions/
// 要先画图，得出思路，然后再写代码，不要急匆匆写

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
		Right: &TreeNode{Val: 2, Left: nil, Right: nil},
	}
	mirrorTree(root)
}
func mirrorTree(root *TreeNode) *TreeNode {
	ans := &TreeNode{}
	ans.Left = root

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		temp := node.Right
		node.Right = node.Left
		node.Left = temp
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return ans.Left
}
