package main

import "fmt"

func sumSlice(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func sumVariadic(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	slice := sumSlice([]int{1, 2, 3, 4, 5})
	variadic := sumVariadic(1, 2, 3, 4, 5)

	fmt.Printf("Slice: %d, Variadic: %d\n", slice, variadic)
}
