package main

import "github.com/jakha/lc/pkg"

func postorderTraversal(root *pkg.TreeNode) []string {
	var (
		vals = make([]string, 0)
	)
	node := root
	rootStack := new(stackTreeNode)
	for node != nil {
		switch {
		case node.Left != nil:
			cutNode := *node
			cutNode.Left = nil
			rootStack.Push(&cutNode)
			node = node.Left
			continue
		case node.Right != nil:
			cutNode := *node
			cutNode.Right = nil
			rootStack.Push(&cutNode)
			node = node.Right
			continue
		}

		vals = append(vals, node.Val)
		node = rootStack.Pop()
	}

	return vals
}
