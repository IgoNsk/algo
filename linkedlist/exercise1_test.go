package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExercise1(t *testing.T) {
	tests := []struct {
		name         string
		givenList    []int
		expectedList []int
	}{
		{
			name:         "Value in the end",
			givenList:    []int{3, 2, 1},
			expectedList: []int{3, 2, 1},
		},
		{
			name:         "Value in the end",
			givenList:    []int{1, 1, 2},
			expectedList: []int{1, 2},
		},
		{
			name:         "Value in the end",
			givenList:    []int{1, 2, 1},
			expectedList: []int{1, 2},
		},
		{
			name:         "Value in the end",
			givenList:    []int{2, 1, 1},
			expectedList: []int{2, 1},
		},
		{
			name:         "Value in the end",
			givenList:    []int{1, 1, 1},
			expectedList: []int{1},
		},
		{
			name:         "Value in the end",
			givenList:    []int{1, 3, 1, 2, 1, 2, 2},
			expectedList: []int{1, 3, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ll := LinkedListFromIntSlice(tc.givenList)
			exercise1(ll)
			actualValue := linkedListToIntSlice(ll)

			assert.Equal(t, tc.expectedList, actualValue)
		})
	}
}

func linkedListToIntSlice(ll *LinkedList) []int {
	res := make([]int, 0)

	for item := ll.First(); item != nil; item = item.Next() {
		res = append(res, item.Data().(int))
	}

	return res
}
