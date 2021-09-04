package pkg

type Queue interface {
	Front() (*TreeNode, int)
	Push(node *TreeNode, level int)
	Len() int
}

func (q *QueueTreeNode) Front() (*TreeNode, int64) {
	result := q.top.value
	level := q.top.level
	q.top = q.top.next
	return result, level
}

func (q *QueueTreeNode) Push(node *TreeNode, level int64) {
	q.size++
	item := &queueTreeNodeItem{
		value: node,
		level: level,
	}
	if q.top == nil {
		q.top = item
		q.tail = q.top
		return
	}
	q.tail.next = item
	q.tail = q.tail.next
}

func (q *QueueTreeNode) Len() int {
	return q.size
}

type queueTreeNodeItem struct {
	value *TreeNode
	level int64
	next  *queueTreeNodeItem
}

type QueueTreeNode struct {
	top  *queueTreeNodeItem
	tail *queueTreeNodeItem
	size int
}

// ========================================================

func NewQueue() Queue {
	return &node{}
}

func (q *node) Front() (*TreeNode, int) {
	if q.top == nil {
		return nil, 0
	}
	result := q.top.current
	level := q.top.level
	q.top = q.top.next
	return result, level
}

func (q *node) Push(d *TreeNode, level int) {
	q.size++
	if q.top == nil {
		q.top = &itemL{
			current: d,
			level:   level,
		}
		q.tail = q.top
		return
	}
	q.tail.next = &itemL{
		current: d,
		level:   level,
	}
	q.tail = q.tail.next
}

func (q *node) Len() int {
	return q.size
}

type itemL struct {
	current *TreeNode
	level   int
	next    *itemL
}

type node struct {
	top  *itemL
	tail *itemL
	size int
}
