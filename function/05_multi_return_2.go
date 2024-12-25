package main

import "fmt"

func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}

	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide2(9, 3)
	fmt.Println(c, success)
	d, success := Divide2(9, 0)
	fmt.Println(d, success)
}
