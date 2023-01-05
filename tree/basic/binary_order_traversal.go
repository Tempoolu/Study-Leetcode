package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main2() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: nil, Right: nil}, Right: &TreeNode{Val: 5, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}},
	}
	fmt.Println(orderTraversal(root))
}

// 前序遍历
func orderTraversal(root *TreeNode) []int {
	res := []int{}
	order(root, &res)
	return res
}

func order(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	// *res = append(*res, node.Val) // 前序遍历（根，左，右）
	if node.Left != nil {
		order(node.Left, res)
	}

	// *res = append(*res, node.Val) // 中序遍历（左，根，右）
	if node.Right != nil {
		order(node.Right, res)
	}

	*res = append(*res, node.Val) // 后序遍历（左，右，根）

	return
}
