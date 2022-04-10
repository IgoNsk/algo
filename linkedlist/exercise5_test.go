package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExercise5(t *testing.T) {
	tests := []struct {
		givenList []int
		expected  bool
	}{
		{
			givenList: []int{1, 2, 3},
			expected:  false,
		},
		{
			givenList: []int{1, 2, 3, 1},
			expected:  false,
		},
		{
			givenList: []int{2, 3, 3, 1},
			expected:  false,
		},
		{
			givenList: []int{1, 2, 1},
			expected:  true,
		},
		{
			givenList: []int{1, 2, 2, 1},
			expected:  true,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			actual := exercise5(LinkedListFromIntSlice(tc.givenList))
			assert.Equal(t, tc.expected, actual)
		})
	}
}
