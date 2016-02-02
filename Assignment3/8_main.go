package main

import "fmt"

func half(i int) (a int, even bool) {

	a = i / 2
	if i%2 == 0 {
		fmt.Printf("%t\n", true)
	} else {
		fmt.Printf("%t\n", false)
	}
	return a, even
}
func main() {
	var num int
	fmt.Println("Enter the number :")
	fmt.Scanf("%d", &num)
	fmt.Println(half(num))
}
