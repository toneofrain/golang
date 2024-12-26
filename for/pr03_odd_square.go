package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	for i := 1; i < 10; i++ {
		if isEven(i) {
			continue
		}
		fmt.Printf("%d * %d = %d\n", i, i, i*i)
	}
}
