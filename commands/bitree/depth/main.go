package main

import (
	"fmt"

	"github.com/jakha/lc/mock"
	"github.com/jakha/lc/pkg"
)

func main() {

	var maxDepthV = 0

	fmt.Println(maxDepth(mock.GetLetterNodes(), maxDepthV))
}

func maxDepth(root *pkg.TreeNode, maxDepthV int) int {
	if root == nil {
		return maxDepthV
	}
	maxDepthV++

	if root.Left == nil && root.Right == nil {
		return maxDepthV
	}

	var (
		leftDepth  int
		rightDepth int
	)
	if root.Left != nil {
		leftDepth = maxDepth(root.Left, maxDepthV)
	}
	if root.Right != nil {
		rightDepth = maxDepth(root.Right, maxDepthV)
	}

	if leftDepth > rightDepth {
		return leftDepth
	}

	return rightDepth
}
