package main

import (
	"fmt"

	"github.com/jakha/lc/mock"
	"github.com/jakha/lc/pkg"
)

func main() {
	levelOrderTraversal(mock.GetLetterNodes())
}

func levelOrderTraversal(root *pkg.TreeNode) {
	cortege := make([][]string, 0)
	nodeQueue := pkg.NewQueue()
	node := root
	level := 0
	for node != nil {
		if len(cortege) < level+1 {
			cortege = append(cortege, []string{})
		}
		cortege[level] = append(cortege[level], node.Val)
		level++
		if node.Left != nil {
			nodeQueue.Push(node.Left, level)
		}
		if node.Right != nil {
			nodeQueue.Push(node.Right, level)
		}
		node, level = nodeQueue.Front()
	}

	fmt.Println(cortege)
}
