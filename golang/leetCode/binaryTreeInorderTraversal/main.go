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

func recursiveinorderTraversal(root *TreeNode) []int {
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
		traverse(node.Left)
		path = append(path, node.Val)
		traverse(node.Right)
	}

	traverse(root)
	return path
}

// 效率很好但破懷 link
func breakInorderTraversal(root *TreeNode) []int {
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
		// 因為為 LVR, 所以按照右中左放入堆疊 (確保左子節點先彈出處理)
		if node.Right != nil {
			stack = append(stack, node.Right)
			node.Right = nil
		}

		if node.Left != nil {
			stack = append(stack, node)
			stack = append(stack, node.Left)
			node.Left = nil
		} else {
			path[pos] = node.Val
			pos++
		}
	}

	return path
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	size := countNodes(root)
	// 預估容量減少記憶體重新分配 (平均節點數約為樹高度相關)
	stack := make([]*TreeNode, 0, size)
	path := make([]int, size)
	pos := 0

	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		path[pos] = curr.Val
		pos++
		curr = curr.Right
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
	list := []any{6, 2, 9, 1, 4, 7, nil, nil, nil, 3, 5, nil, 8}
	root := setBinaryTree(list)

	var path []int
	path = recursiveinorderTraversal(root)
	print(path)

	path = []int{}
	path = inorderTraversal(root)
	print(path)
}
