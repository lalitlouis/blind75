package main

func main() {
	TestRecursion()
}

func TestRecursion() {

}

func Recursion(num int) int {

	if num <= 1 {
		return 1
	}

	return num * Recursion(num-1)

}
