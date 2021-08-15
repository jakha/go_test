package pkg

type Queue interface {
	Front() (*TreeNode, int64)
	Push(node *TreeNode, level int64)
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

func (q *node) Front() string {
	result := q.top.value
	q.top = q.top.next
	return result
}

func (q *node) Push(d string) {
	q.size++
	if q.top == nil {
		q.top = &itemS{
			value: d,
		}
		q.tail = q.top
		return
	}
	q.tail.next = &itemS{
		value: d,
	}
	q.tail = q.tail.next
}

func (q *node) Len() int {
	return q.size
}

type itemS struct {
	value string
	next  *itemS
}

type node struct {
	top  *itemS
	tail *itemS
	size int
}
