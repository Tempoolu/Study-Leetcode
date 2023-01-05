package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 计算到叶子节点的所有路径，简单，自己写出来，和路径有关系，还和扩展参数有关系
// https://leetcode.cn/problems/binary-tree-paths/comments/
// 同样类型的题目还有：https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/

func main() {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 4, Left: nil, Right: nil}, Right: nil},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: nil, Right: nil}, Right: nil},
	}

	fmt.Println(binaryTreePaths(root))

}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := [][]int{}

	var dfs func(node *TreeNode, path []int)
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, append(path, node.Val))
			return
		}
		dfs(node.Left, append(path, node.Val))
		dfs(node.Right, append(path, node.Val))
	}
	dfs(root, []int{})

	str := []string{}
	for _, r := range res {
		temp := ""
		for _, rr := range r {
			temp += strconv.Itoa(rr) + "->"
		}
		if strings.HasSuffix(temp, "->") {
			temp = strings.TrimRight(temp, "->")
		}
		str = append(str, temp)
	}
	return str
}
