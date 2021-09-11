package main

import (
	"github.com/jakha/lc/mock"
	"github.com/jakha/lc/pkg"
)

func main() {
	isSymmetric(mock.GetLetterNodes())
}

func isSymmetric(root *pkg.TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}

	leftBranch := lBranchNodes(root.Left)
	rightBranch := rBranchNodes(root.Right)
	if len(leftBranch) != len(rightBranch) {
		return false
	}

	for i, lv := range leftBranch {
		if lv != rightBranch[i] {
			return false
		}
	}

	return true
}

func lBranchNodes(root *pkg.TreeNode) []string {
	var nodeOrders = []string{root.Val}

	if root.Left != nil {
		nodeOrders = append(nodeOrders, lBranchNodes(root.Left)...)
	} else {
		nodeOrders = append(nodeOrders, "")
	}

	if root.Right != nil {
		nodeOrders = append(nodeOrders, lBranchNodes(root.Right)...)
	} else {
		nodeOrders = append(nodeOrders, "")
	}

	return nodeOrders
}

func rBranchNodes(root *pkg.TreeNode) []string {
	var nodeOrders = []string{root.Val}

	if root.Right != nil {
		nodeOrders = append(nodeOrders, rBranchNodes(root.Right)...)
	} else {
		nodeOrders = append(nodeOrders, "")
	}

	if root.Left != nil {
		nodeOrders = append(nodeOrders, rBranchNodes(root.Left)...)
	} else {
		nodeOrders = append(nodeOrders, "")
	}

	return nodeOrders
}
