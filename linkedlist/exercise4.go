package linkedlist

// Два числа хранятся в виде двух связанных списков, в которых каждый узел представляет собой 1 разряд. Все цифры
// хранятся в обратном порядке, при этом младший разряд хранится в начале списка.
// Напишите функцию, которая суммирует два числа и возвращает результат в виде нового списка.
//
// Сложность решения O(max(N,M)) - где N и M это кол-во элементов в каждом списке
func exercise4(ll1, ll2 *LinkedList) *LinkedList {
	acc := NewEmptyLinkedList()

	var overflowValue bool
	for node1, node2 := ll1.First(), ll2.First(); node1 != nil || node2 != nil; {
		var v int
		if node1 != nil {
			v += node1.Data().(int)
			node1 = node1.Next()
		}
		if node2 != nil {
			v += node2.Data().(int)
			node2 = node2.Next()
		}

		if overflowValue {
			v += 1
		}
		if v >= 10 {
			v %= 10
			overflowValue = true
		} else {
			overflowValue = false
		}

		acc.AppendToTail(v)
	}

	if overflowValue {
		acc.AppendToTail(1)
	}

	return acc
}
