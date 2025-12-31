package main

import "fmt"

func main() {
	result := func(word1, word2 string) string {
		return word1 + " " + word2
	}("hello", "world")
	fmt.Println(result)
}
