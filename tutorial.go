package main

import (
	"fmt"
	"math"
)

// if something has an area and returns a float64, it is type shape
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func GetArea(s shape) float64 {
	return s.area()
}

// entry point into our app. Will be called when we run our go program
func main() {
	c1 := circle{4.5}
	r1 := rect{5, 7}

	// when calling shapes, we can only use .area
	shapes := []shape{&c1, &r1}

	for _, shape := range shapes {
		fmt.Println(GetArea(shape))
	}
}

// variable - a way of storing and accessing information
// function - a block of reusable code

// interface - a way of looking at a set of related objects or types
