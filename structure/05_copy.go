package main

import "fmt"

type Student struct {
	Age   int
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("Age: %d, No: %d, Score: %f", s.Age, s.No, s.Score)
}

func main() {
	var student = Student{15, 23, 88.2}

	student2 := student

	PrintStudent(student2)
}
