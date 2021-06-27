package main

import (
	"context"
	"fmt"
	"time"
)

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	/**
	 * Definition for a binary tree node.
	 * type TreeNode struct {
	 *     Val int
	 *     Left *TreeNode
	 *     Right *TreeNode
	 * }
	 */

	//s := postorderTraversal(getLetterNodes())

	//t := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	//j := 0
	//for _, v := range s {
	//	if v != t[j] {
	//		panic("wrong")
	//	}
	//	j++
	//}

	levelOrderTraversal(
		getLetterNodes())

	// ctx := context.Background()
	// ctx, _ = context.WithTimeout(ctx, time.Second)
	// defer cancel()

	// sleepAndTalk(ctx, 5*ti
	//me.Second, "hello")

	// time.Sleep(10 * time.Second)
	//	cancel()go

}

func preorderTraversal(root *TreeNode) []string {
	vals := make([]string, 0)
	rights := make([]*TreeNode, 0)
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

func postorderTraversal(root *TreeNode) []string {
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

func levelOrderTraversal(root *TreeNode) {
	queue := new(queueTreeNode)
	// result := make([][]string, 0) 
	node := root
	resultSub := make([]string, 0)
	var level int64 = 0
	iteration := 0
	for node != nil {
		if node.Left != nil {
			nLevel := level + 1
			queue.Push(node.Left, nLevel)
		}
		if node.Right != nil {
			nLevel := level + 1
			queue.Push(node.Right, nLevel)
		}

		node, level = queue.Front()
		iteration++
	}
}

func sleepAndTalk(ctx context.Context, dur time.Duration, txt string) {
	select {
	case <-time.After(dur):
		fmt.Println(txt)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

//package main
//
//import (
//"fmt"
//"sync"
//"strconv"
//)

//func main() {
//	var wc sync.WaitGroup
//	m := make(chan string, 5)
//	fff := sync.Mutex{}
//	for i := 0; i < 5; i++ {
//		wc.Add(1)
//		go func(mm chan<- string, i int, group *sync.WaitGroup) {
//			defer wc.Done()
//			fff.Lock()
//			mm <- fmt.Sprintf("Goroutine %s", strconv.Itoa(i))
//			fff.Unlock()
//		}(m, i, &wc)
//	}
//
//	wc.Wait()
//	close(m)
//
//	for q := range m {
//		fmt.Println(q)
//	}
//
//}
