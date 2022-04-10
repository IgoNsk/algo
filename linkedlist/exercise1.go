package linkedlist

// Напишите код для удаления дубликатов из несортированного связанного списка.
// Использовать временный буфер - запрещено.
//
// Оценка сложности - O(N^2), расход памяти - константный
func exercise1(ll *LinkedList) {
	for item := ll.First(); item != nil; item = item.Next() {
		for comparedItem := item.Next(); comparedItem != nil; comparedItem = comparedItem.Next() {
			if item.Data() == comparedItem.Data() {
				ll.RemoveNode(comparedItem)
			}
		}
	}
}
