package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExercise3(t *testing.T) {
	tests := []struct {
		name          string
		givenList     []int
		valueToDelete int
		expectedList  []int
	}{
		{
			givenList:     []int{1, 2, 3, 4},
			valueToDelete: 2,
			expectedList:  []int{1, 3, 4},
		},
		{
			givenList:     []int{1, 2, 3, 4},
			valueToDelete: 3,
			expectedList:  []int{1, 2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ll := LinkedListFromIntSlice(tc.givenList)

			exercise3(ll.FindNodeWithValue(tc.valueToDelete))
			actualValue := linkedListToIntSlice(ll)

			assert.Equal(t, tc.expectedList, actualValue)
		})
	}
}
