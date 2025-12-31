package main

import "fmt"

func divide(num, div int) (res, rem int) {
	res = num / div
	rem = num % div
	return res, rem
}

func main() {
	res, rem := divide(3, 2)
	fmt.Printf("Result: %d, Remainder: %d\n", res, rem)
}
