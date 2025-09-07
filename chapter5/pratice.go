package main

import "fmt"

func Calc(a int, b int) int {
	return a * b
}

func Fibo(n int) int {
	memo := make(map[int]int)
	return fiboWithMemo(n, memo)
}

func fiboWithMemo(n int, memo map[int]int) int {
	if val, exists := memo[n]; exists {
		fmt.Printf("Fibo(%d): 메모에서 값 %d 반환\n", n, val)
		return val
	}
	if n < 2 {
		fmt.Printf("Fibo(%d): 기저 사례, 값 %d 반환\n", n, n)
		memo[n] = n
		return n
	}
	fmt.Printf("Fibo(%d): 계산 중, Fibo(%d) + Fibo(%d)\n", n, n-2, n-1)
	result := fiboWithMemo(n-2, memo) + fiboWithMemo(n-1, memo)
	memo[n] = result
	fmt.Printf("Fibo(%d): 계산 완료, 값 %d 저장 및 반환\n", n, result)
	return result
}

func main() {
	result := Calc(10, 10)
	fmt.Println(result)
	fmt.Println(Fibo(9))
}
