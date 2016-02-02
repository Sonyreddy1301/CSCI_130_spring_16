package main

import "fmt"

func maximum(numbers ...int) int {
	var a int
	for _, b := range numbers {
		if b > a {
			a = b
		}
	}
	return a
}
func main() {
	maxvalue := maximum(13, 2, 20, 1, 8000, 40, 60)
	fmt.Println(maxvalue)
}
