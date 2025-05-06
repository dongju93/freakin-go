package main

import "fmt"

func main() {
	// Integer Overflow
	var x int8 = 127 // 8비트 부호가 있는 정수 최댓값

	/*
		-128(1000 0000) ~ 127(0111 1111)
		127 에서 1을 더하여 Integer Overflow 발생: -128
	*/
	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t= %4d, %08b\n", x, uint8(x))
	fmt.Printf("x + 1\t= %4d, %08b\n", x+1, uint8(x+1))
	fmt.Printf("x + 2\t= %4d, %08b\n", x+2, uint8(x+2))
	fmt.Printf("x + 3\t= %4d, %08b\n", x+3, uint8(x+3))

	var y int8 = -128
	fmt.Printf("%d < %d - 1: %v\n", y, y, y < y-1)
	fmt.Printf("y\t= %4d, %08b\n", y, uint8(y))
	fmt.Printf("y - 1\t= %4d, %08b\n", y-1, uint8(y-1))

	// Integer Underflow
	fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1)
	fmt.Printf("y - 1\t= %4d, %08b\n", y-1, uint8(y-1))
}
