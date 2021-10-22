package main

import (
	"fmt"
)

func main() {
	var n int
	var start int
	fmt.Println("введите n, start")
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &start)
	fmt.Println("Ответ: ", xorOperation(n, start))
}

func xorOperation(n int, start int) int {
	array := make([]int, n)
	var xor int
	xor = 0
	for i := 0; i < n; i++ {
		array[i] = start + 2*i
		xor = xor ^ array[i]
	}
	return xor
}
