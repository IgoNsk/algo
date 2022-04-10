package linkedlist

import "errors"

// Реализуйте алгоритм для нахождения в односвязном списке k-го элемента с конца
//
// Используем "бегунок" - позицию в списке, от которой просматриваем список вперед, заглядываем до тех пор, пока не продвинемся
// вперед на необходимое смещение. Когда до него доходим - проверяем, является ли следующий элемент пустым. Если да - мы дошли
// до конца, и тогда элемент на котором находится "бегунок", считается находящимся на расстоянии K от конца списка.
//
// Сложность алгоритма в худшем случае O(n*K) - т.к. мы пройдем список полностью, но при этом не более чем K раз для каждого элемента
func exercise2(ll *LinkedList, k int) (interface{}, error) {
	if ll.First() == nil {
		return nil, errors.New("empty linked list")
	}

	if k <= 0 {
		return nil, errors.New("K should be positive and not empty integer value")
	}

MainLoop:
	for cursor := ll.First(); cursor != nil; cursor = cursor.Next() {

		comparedItem := cursor
		for i := 1; i < k; i++ {
			// Оптимизация. Если мы дошли до конца списка, и при этом не смогли уйти вперед на заданное число позиций,
			// то остановиться можно прямо сейчас. Тем самым сделаем сложность O(n)
			if comparedItem == nil {
				break MainLoop
			}
			comparedItem = comparedItem.Next()
		}
		if comparedItem.Next() == nil {
			return cursor.Data(), nil
		}
	}

	return nil, nil
}
