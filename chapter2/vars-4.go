package main

import "fmt"

func main() {
	var a float32 = 10.1234567 // float32 의 최대 소수점 자리수는 7자리
	var b float64 = 10.1234567 // float64 의 최대 소수점 자리수는 15자리
	var c float32 = a * a      // float32 의 최대 소수점 자리수 초과 시
	var d float64 = b * b      // float64 의 최대 소수점 자리수 초과 시

	fmt.Println(a, b)
	fmt.Println(c, d)
}
