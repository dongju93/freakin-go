package main

import (
	"fmt"
)

// 첫글자가 대문자인 함수는 패키지 외부로 공개되는 함수
func Add(a int, b int) int {
	return a + b
}

func main() {
	c := Add(3, 6)
	fmt.Println((c))
}
