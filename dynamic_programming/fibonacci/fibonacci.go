package main

import "fmt"

func main() {
	TestFibonacci()
}

func TestFibonacci() {
	testCases := []struct {
		name     string
		input    int
		expected int
		desc     string
	}{
		{"base case 0", 0, 0, "F(0) = 0"},
		{"base case 1", 1, 1, "F(1) = 1"},
		{"small values", 2, 1, "F(2) = F(1) + F(0) = 1"},
		{"fibonacci 3", 3, 2, "F(3) = F(2) + F(1) = 2"},
		{"fibonacci 4", 4, 3, "F(4) = F(3) + F(2) = 3"},
		{"fibonacci 5", 5, 5, "F(5) = F(4) + F(3) = 5"},
		{"fibonacci 6", 6, 8, "F(6) = F(5) + F(4) = 8"},
		{"fibonacci 7", 7, 13, "F(7) = F(6) + F(5) = 13"},
		{"fibonacci 8", 8, 21, "F(8) = F(7) + F(6) = 21"},
		{"fibonacci 10", 10, 55, "F(10) = 55"},
	}

	for _, tc := range testCases {
		memo := make(map[int]int, 0)
		result := Fibonacci(tc.input, memo)

		fmt.Printf("Test: %s\n", tc.desc)
		fmt.Printf("Input: %d, Expected: %d, Got: %d\n", tc.input, tc.expected, result)

		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Printf("FAIL\n\n")
		}
	}
}

// 0 1 1 2 3 5 8
func Fibonacci(num int, memo map[int]int) int {

	if num <= 1 {
		return num
	}

	if value, ok := memo[num]; ok {
		return value
	}
	memo[num] = Fibonacci(num-2, memo) + Fibonacci(num-1, memo)

	return memo[num]
}
