package main

import "fmt"

// entry point into our app. Will be called when we run our go program
func main() {
	// under the hood, Go will create an array and then a slice of the array that encompases the entire array
	a := []int{1, 2, 3}

	a = append(a, 10)

	fmt.Println(a)
	
}

// variable - a way of storing and accessing information