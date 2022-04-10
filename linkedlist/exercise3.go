package linkedlist

// Реализуйте алгоритм, удаляющий узел из середины односвязного списка. Доступ предоставляется только к этому узлу
//
// Сложность O(1)
func exercise3(n *Node) {
	n.data = n.next.data
	if n.next.next != nil {
		n.next.next.prev = n
	}
	n.next = n.next.next
}
