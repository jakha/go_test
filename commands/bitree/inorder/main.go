package main

func inorderTraversal(root *TreeNode) []string {
	var vals = make([]string, 0)
	node := root
	nodesStack := new(stackTreeNode)
	fromStack := false
	for node != nil {
		if node.Left != nil && !fromStack {
			nodesStack.Push(node)
			node = node.Left
			continue
		}
		vals = append(vals, node.Val)
		if node.Right != nil {
			fromStack = false
			node = node.Right
			continue
		}

		fromStack = true
		node = nodesStack.Pop()
	}

	return vals
}
