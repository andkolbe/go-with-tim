package main

import "fmt"

func add(x, y int) (z1, z2 int) {
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	fmt.Println("before return")
	return
}

// entry point into our app. Will be called when we run our go program
func main() {
	ans1, ans2 := add(4, 2)
	fmt.Println(ans1, ans2)

}

// variable - a way of storing and accessing information
// function - a block of reusable code