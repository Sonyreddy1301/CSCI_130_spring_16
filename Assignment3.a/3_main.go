package main

import "fmt"

func main() {
	var name string
	fmt.Println("Hello user may I know your name:")
	fmt.Scanf("%s", &name)
	message := "Hi" + name
	fmt.Println(message)
}
