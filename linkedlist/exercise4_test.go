package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExercise4(t *testing.T) {
	tests := []struct {
		name         string
		givenList1   []int
		givenList2   []int
		expectedList []int
	}{
		{
			name:         "Строки одинаковой длинной без переполнения",
			givenList1:   []int{2, 0, 2},
			givenList2:   []int{5, 9, 2},
			expectedList: []int{7, 9, 4},
		},
		{
			name:         "Строки одинаковой длинной с переполнением в начале",
			givenList1:   []int{7, 1, 6},
			givenList2:   []int{5, 9, 2},
			expectedList: []int{2, 1, 9},
		},
		{
			name:         "Строки одинаковой длинной с переполнением в середине",
			givenList1:   []int{7, 1, 6},
			givenList2:   []int{1, 9, 2},
			expectedList: []int{8, 0, 9},
		},
		{
			name:         "Строки одинаковой длинной с переполнением в конце",
			givenList1:   []int{7, 1, 6},
			givenList2:   []int{1, 8, 5},
			expectedList: []int{8, 9, 1, 1},
		},
		{
			name:         "Строки разной длинны - первая короче",
			givenList1:   []int{7},
			givenList2:   []int{1, 8, 5},
			expectedList: []int{8, 8, 5},
		},
		{
			name:         "Строки разной длинны - вторая короче",
			givenList1:   []int{1, 8, 5},
			givenList2:   []int{7, 2},
			expectedList: []int{8, 0, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := exercise4(LinkedListFromIntSlice(tc.givenList1), LinkedListFromIntSlice(tc.givenList2))
			actualValue := linkedListToIntSlice(actualResult)

			assert.Equal(t, tc.expectedList, actualValue)
		})
	}
}
