package array

import (
	"reflect"
	"testing"
)

func verifySums(t *testing.T, result, expect []int) {
	t.Helper()
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result '%d', expected '%d'", result, expect)
	}
}

func TestSum(t *testing.T) {
	

	t.Run("any size collection", func(t *testing.T) {
		// Arrange
		numbers := []int{1, 2, 3}
		expect := 6

		// Act
		result := Sum(numbers)

		// Assert
		if expect != result {
			t.Errorf("result '%d', expected '%d', input '%v'", result, expect, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	// Arrange
	expect := []int{3, 9}

	// Act
	result := SumAll([]int{1, 2}, []int{0, 9})

	// Assert
	verifySums(t, result, expect)
}

func TestSumRest(t *testing.T) {
	t.Run("sum the rest of some slices", func(t *testing.T) {
		// Arrange
		expect := []int{2, 9}

		// Act
		result := SumRest([]int{1, 2}, []int{0, 9})

		// Assert
		verifySums(t, result, expect)
	})

	t.Run("sum empty slices in a safe way", func(t *testing.T) {
		// Arrange
		expect := []int{0, 9}

		// Act
		result := SumRest([]int{}, []int{0, 9})

		// Assert
		verifySums(t, result, expect)
	})
}
