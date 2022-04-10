package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExercise2_Positive(t *testing.T) {
	tests := []struct {
		name          string
		givenList     []int
		givenK        int
		expectedValue int
	}{
		{
			name:          "Value in the end",
			givenList:     []int{33, 22, 11},
			givenK:        1,
			expectedValue: 11,
		},
		{
			name:          "Value in the middle",
			givenList:     []int{33, 22, 11},
			givenK:        2,
			expectedValue: 22,
		},
		{
			name:          "Value in the beginning",
			givenList:     []int{33, 22, 11},
			givenK:        3,
			expectedValue: 33,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualValue, err := exercise2(LinkedListFromIntSlice(tc.givenList), tc.givenK)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedValue, actualValue)
		})
	}

}

func TestExercise2_IncorrectNegativeK(t *testing.T) {
	givenK := -1
	_, err := exercise2(LinkedListFromIntSlice([]int{33, 22, 11}), givenK)

	assert.Error(t, err)
}

func TestExercise2_IncorrectZeroK(t *testing.T) {
	givenK := 0
	_, err := exercise2(LinkedListFromIntSlice([]int{33, 22, 11}), givenK)

	assert.Error(t, err)
}

func TestExercise2_KMoreThanList(t *testing.T) {
	givenK := 10
	actualValue, err := exercise2(LinkedListFromIntSlice([]int{33, 22, 11}), givenK)

	assert.NoError(t, err)
	assert.Nil(t, actualValue)
}

func LinkedListFromIntSlice(s []int) *LinkedList {
	var ll *LinkedList
	for _, i := range s {
		if ll == nil {
			ll = NewLinkedList(i)
		} else {
			ll.AppendToTail(i)
		}
	}
	return ll
}
