package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val:      2,
				Children: []*Node{{Val: 5, Children: nil}, {Val: 6, Children: nil}},
			},
			{
				Val:      3,
				Children: []*Node{{Val: 7, Children: nil}},
			},
			{
				Val:      4,
				Children: []*Node{{Val: 8, Children: nil}, {Val: 9, Children: []*Node{{Val: 10, Children: nil}}}},
			},
		},
	}
	fmt.Println(nTreeOrderTraversal(root))
}

func nTreeOrderTraversal(root *Node) []int {
	arr := []int{}
	nTreeOrder(root, &arr)
	return arr
}

func nTreeOrder(node *Node, res *[]int) {
	if node == nil {
		return
	}
	// *res = append(*res, node.Val)    // 前序遍历 [1 2 5 6 3 7 4 8 9 10]
	for _, child := range node.Children {
		nTreeOrder(child, res)
	}
	*res = append(*res, node.Val) // 后序遍历 [5 6 2 7 3 8 10 9 4 1]
}
