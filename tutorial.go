package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// a method on the Student struct
// needs to be a pointer because we are modifying the student
func (s *Student) setAge(age int) {
	s.age = age
	fmt.Println(age)
}

// doesn't nedd to be a pointer because we aren't changing anything about the student
func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	curMax := 0
	for _, grade := range s.grades {
		if grade > curMax {
			curMax = grade
		}
	}
	return curMax
}

// entry point into our app. Will be called when we run our go program
func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	fmt.Println(s1.age)
	s1.setAge(7)

	fmt.Println(s1.getAverageGrade())
	fmt.Println(s1.getMaxGrade())

}

// variable - a way of storing and accessing information
// function - a block of reusable code
