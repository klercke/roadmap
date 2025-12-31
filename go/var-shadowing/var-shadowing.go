package main

import "fmt"

var word = "hello"

func main() {
	var word = "world"
	fmt.Println("inside main(): ", word)
	printOuter()
}

func printOuter()  {
	fmt.Println("package level: ", word)
}
