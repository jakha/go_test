package pkg

type StackTreeNodeItem struct {
	value *TreeNode
	next  *StackTreeNodeItem
}

type StackTreeNode struct {
	top  *StackTreeNodeItem
	size int
}

func (s *StackTreeNode) Pop() *TreeNode {
	if s.Len() > 0 {
		topStackVal := s.top.value
		s.top = s.top.next
		s.size--
		return topStackVal
	}
	return nil
}

func (s *StackTreeNode) Push(treeNode *TreeNode) {
	s.top = &StackTreeNodeItem{
		value: treeNode,
		next:  s.top,
	}
	s.size++
}

func (s *StackTreeNode) Len() int {
	return s.size
}

type Stack interface {
	Pop() int
	Push(val int)
	Len() int
}

type item struct {
	value int
	next  *item
}

type stack struct {
	top  *item
	size int
}

func (s *stack) Pop() int {
	if s.Len() > 0 {
		topStackVal := s.top.value
		s.top = s.top.next
		s.size--
		return topStackVal
	}
	return 0
}

func (s *stack) Push(val int) {
	s.top = &item{
		value: val,
		next:  s.top,
	}
	s.size++
}

func (s *stack) Len() int {
	return s.size
}
