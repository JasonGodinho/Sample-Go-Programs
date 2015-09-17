package main_test

import "testing"

var expected = map[int]int{
	0:  0,
	1:  1,
	2:  1,
	3:  2,
	4:  3,
	5:  5,
	6:  8,
	7:  13,
	8:  21,
	9:  34,
	10: 55,
	11: 89,
	12: 144,
	13: 233,
	14: 377,
	15: 610,
	16: 987,
	17: 1597,
	18: 2584,
	19: 4181,
	20: 6765,
}

func TestRecursuveFibonacci(t *testing.T) {

	var i int
	for itemCount := int(len(expected)); i < itemCount; i++ {
		if result := RecursuveFibonacci(i); result != expected[i] {
			t.Errorf("RecursuveFibonacci(%d) returned %d, expected %d", i, result, expected[i])
		}
	}
}

func RecursuveFibonacci(fibonacci_num int) int {
	if fibonacci_num == 1 {
		return 1
	} else if fibonacci_num == 0 {
		return 0
	} else {
		return RecursuveFibonacci(fibonacci_num-1) + RecursuveFibonacci(fibonacci_num-2)
	}
}
