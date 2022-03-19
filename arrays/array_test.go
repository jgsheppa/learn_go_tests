package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)
	})
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3})
	fmt.Println(sum)
	// Output: 6
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSumAll() {
	sum := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(sum)
	// Output: [3 9]
}


func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	sum := SumAllTails([]int{1, 2}, []int{0, 9})
	fmt.Println(sum)
	// Output: [2 9]
}