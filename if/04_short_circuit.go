package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 increased")
	}
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 increased")
	}

	fmt.Println("cnt:", cnt)
}
