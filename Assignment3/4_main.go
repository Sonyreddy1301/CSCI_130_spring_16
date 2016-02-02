package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Printf("Enter the larger number:")
	fmt.Scanf("%d\n", &x)
	fmt.Printf("Enter the smaller number:")
	fmt.Scanf("%d\n", &y)
	remainder := x % y
	fmt.Println("The remainder is:", remainder)

}
