package linkedlist

type LinkedList struct {
	firstItem *Node
	lastItem  *Node
}

func NewLinkedList(data interface{}) *LinkedList {
	n := &Node{data, nil, nil}
	return &LinkedList{n, n}
}

func NewEmptyLinkedList() *LinkedList {
	return &LinkedList{nil, nil}
}

func (r LinkedList) First() *Node {
	return r.firstItem
}

func (r LinkedList) Last() *Node {
	return r.lastItem
}

func (r *LinkedList) AppendToHead(data interface{}) {
	n := Node{
		data: data,
		prev: nil,
		next: r.firstItem,
	}

	if r.firstItem != nil {
		r.firstItem.prev = &n
	} else {
		r.firstItem = &n
	}

	r.firstItem = &n
}

func (r *LinkedList) AppendToTail(data interface{}) {
	n := Node{
		data: data,
		prev: r.lastItem,
		next: nil,
	}

	if r.lastItem != nil {
		r.lastItem.next = &n
	} else {
		r.firstItem = &n
	}

	r.lastItem = &n
}

func (r *LinkedList) FindNodeWithValue(data interface{}) *Node {
	for node := r.First(); node != nil; node = node.Next() {
		if node.Data() == data {
			return node
		}
	}

	return nil
}

func (r *LinkedList) RemoveNode(n *Node) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		r.firstItem = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		r.lastItem = n.prev
	}
}

type Node struct {
	data interface{}

	next *Node
	prev *Node
}

func (r Node) Data() interface{} {
	return r.data
}

func (r Node) Next() *Node {
	return r.next
}

func (r Node) Prev() *Node {
	return r.prev
}
