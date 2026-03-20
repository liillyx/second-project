package main

import "fmt"

func main() {
	fmt.Println("App started")
	result := Process("Hello", "Word")
	fmt.Println("Result:", result)

	reversed := Reverse("Hello")
	fmt.Println("Reversed:", reversed)
}
