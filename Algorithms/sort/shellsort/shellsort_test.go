package shellsort

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{0}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}
	for i, tt := range tests {
		ShellSort(tt.input)
		if got := isSorted(tt.input); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}

}

func isSorted(elements []int) bool {
	l := len(elements)
	for i := 0; i < l-1; i++ {
		if elements[i] > elements[i+1] {
			return false
		}
	}
	return true
}
