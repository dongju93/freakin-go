package main

import "fmt"

const PI = 3.14              // 타입 미지정
const FloatPI float64 = 3.14 // 타입 지정

func main() {
	var a int = PI * 100      // 타입이 지정되지 않아 314 (정수)로 변환되어 저장, 변수에 복사될 때 타입 결정
	var b int = FloatPI * 100 // 타입이 지정되어 314.0 (실수)로 계산되기 때문에 314 (정수) 로 저장 불가능

	fmt.Println(a)
	fmt.Println(b)
}
