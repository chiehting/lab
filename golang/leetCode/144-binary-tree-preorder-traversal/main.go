package main

import (
	"fmt"
)

// TreeNode for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func print(message ...any) {
	fmt.Println(message...)
}

// helper: 只算節點數，回傳 int
func countNodes(n *TreeNode) int {
	if n == nil {
		return 0
	}
	return 1 + countNodes(n.Left) + countNodes(n.Right)
}

func recursivePreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	size := countNodes(root)
	path := make([]int, 0, size)
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return path
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	size := countNodes(root)
	// 預估容量減少記憶體重新分配 (平均節點數約為樹高度相關)
	stack := make([]*TreeNode, 0, size)
	stack = append(stack, root)
	path := make([]int, size)
	pos := 0

	for len(stack) > 0 {
		// 彈出堆疊頂部節點
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		path[pos] = node.Val
		pos++

		// 先壓入右子節點，再壓入左子節點 (確保左子節點先彈出處理)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return path
}

func setBinaryTree(list []any) *TreeNode {
	if len(list) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(list))
	for i, val := range list {
		if val != nil {
			nodes[i] = &TreeNode{Val: val.(int)}
		}
	}

	n := len(nodes)
	for i, node := range nodes {
		if node == nil {
			continue
		}

		leftIdx, rightIdx := 2*i+1, 2*i+2
		if leftIdx < n {
			nodes[i].Left = nodes[leftIdx]
		}
		if rightIdx < n {
			nodes[i].Right = nodes[rightIdx]
		}
	}
	return nodes[0]
}

func main() {
	// list := []any{1, nil, 2, nil, nil, 3}
	list := []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, nil, nil, 9}
	root := setBinaryTree(list)

	var path []int
	path = recursivePreorderTraversal(root)
	print(path)

	path = []int{}
	path = preorderTraversal(root)
	print(path)
}
