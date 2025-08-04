package main

import (
	"fmt"
)

// 반환 타입에 변수명을 지정 시 return 에 명시적으로 선언하지 않더라고 값을 반환
func Divide(a, b int) (result int, success bool) {
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
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
