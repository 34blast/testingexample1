package testingexample1

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("[1,2,3,4,5]", testSumFunc([]int{1, 2, 3, 4, 5}, 15))
	t.Run("[1,2,3,4,-5]", testSumFunc([]int{1, 2, 3, 4, -5}, 5))
}

func testSumFunc(numbers []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := Sum(numbers)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual))
		}
	}
}

func catchAll() {
	fmt.Println("callAll() : entry")
}

func Sum(numbers []int) int {
	sum := 0
	for n := range numbers {
		sum += numbers[n]
	}
	return sum
}

func ExampleSum() {
  numbers := []int{5, 5, 5}
  fmt.Println(Sum(numbers))
  // Output:
  // 15
}

func BenchmarkSumMe(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
