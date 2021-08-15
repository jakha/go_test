package main

import "github.com/jakha/lc/pkg"

func preorderTraversal(root *pkg.TreeNode) []string {
	vals := make([]string, 0)
	rights := make([]*pkg.TreeNode, 0)
	node := root
	for node != nil {
		// go to left
		vals = append(vals, node.Val)
		if node.Left != nil {
			if node.Right != nil {
				rights = append(rights, node.Right)
			}
			node = node.Left
			continue
		}
		if node.Right != nil {
			node = node.Right
			continue
		}

		rightsLen := len(rights) - 1
		if rightsLen == -1 {
			node = nil
			continue
		}
		node = rights[rightsLen]
		rights = rights[:rightsLen]
	}

	return vals
}
