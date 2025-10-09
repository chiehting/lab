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

func recursivePostorderTraversal(root *TreeNode) []int {
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
		traverse(node.Right)
		path = append(path, node.Val)
	}

	traverse(root)
	return path
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	size := countNodes(root)
	// 預估容量減少記憶體重新分配 (平均節點數約為樹高度相關)
	stack := make([]*TreeNode, 0, size)
	path := make([]int, size)
	pos := 0
	var temp *TreeNode

	curr := root
	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {

			temp = stack[len(stack)-1].Right

			if temp == nil {
				temp = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				path[pos] = temp.Val
				pos++

				for len(stack) > 0 && stack[len(stack)-1].Right == temp {
					temp = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					path[pos] = temp.Val
					pos++
				}
			} else {
				curr = temp
			}
		}
	}

	return path
}

func breakPostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	size := countNodes(root)
	// 預估容量減少記憶體重新分配 (平均節點數約為樹高度相關)
	stack := make([]*TreeNode, 0, size)
	path := make([]int, size)
	pos := 0
	var temp *TreeNode

	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
			stack[len(stack)-1].Left = nil
		}

		temp = stack[len(stack)-1]

		if temp == nil {
			continue
		}

		if temp.Left == nil && temp.Right == nil {
			path[pos] = temp.Val
			pos++
			stack = stack[:len(stack)-1]
		}

		if temp.Right != nil {
			curr = temp.Right
			stack[len(stack)-1].Right = nil
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
	list := []any{9, 5, 8, 1, 4, 7, nil, nil, nil, 2, 3, nil, 6}
	root := setBinaryTree(list)

	var path []int
	path = recursivePostorderTraversal(root)
	print(path)

	path = []int{}
	path = postorderTraversal(root)
	print(path)

	path = []int{}
	path = breakPostorderTraversal(root)
	print(path)
}
