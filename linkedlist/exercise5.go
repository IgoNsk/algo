package linkedlist

// Напишите функцию, которая проверяет, является ли связанный список пaлиндромом
//
// Сложность O(N/2)
func exercise5(ll *LinkedList) bool {
	// Идем одновременно с начала и конца списка - сравниваем элементы, и в случае расхождения - это точно не палиндром.
	for fromBeginning, fromEnd := ll.First(), ll.Last(); fromBeginning != fromEnd; fromBeginning, fromEnd = fromBeginning.Next(), fromEnd.Prev() {
		if fromBeginning.Data() != fromEnd.Data() {
			return false
		}

		// Если размерность списка не четная, то следующий шаг приведет нас к расхождению. Если так - то средний элемент
		// не надо проверять - он будет находиться ровно посередине
		if fromBeginning.Next() == fromEnd || fromBeginning == fromEnd.Prev() {
			return true
		}
	}

	return true
}
