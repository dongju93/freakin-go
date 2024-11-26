package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int64 = 10   // 기본적인 변수 선언
	var b int          // 타입 선언만 함, 기본 값 0
	var c = 3.123      // 타입 생략, 대입 값으로 타입 추론
	d := "T bal C ya?" // var 키워드와 타입 생략(walrus 사용), 대입 값으로 타입 추론

	fmt.Println(
		" vars a type:", reflect.TypeOf(a), "and value is:", a, "\n",
		"vars b type:", reflect.TypeOf(b), "and value is:", b, "\n",
		"vars c type:", reflect.TypeOf(c), "and value is:", c, "\n",
		"vars d type:", reflect.TypeOf(d), "and value is:", d, "\n",
		"print null in go is", nil,
	)

	/*
		강타입 언어로 아래와 같은 연산 시 에러 발생
		chapter2/vars-3.go:23:10: invalid operation: a + c (mismatched types int64 and float64)
		chapter2/vars-3.go:24:14: cannot use c (variable of type float64) as int value in variable declaration
		chapter2/vars-3.go:26:7: invalid operation: a * b (mismatched types int64 and int)
	*/
	// var e = a + c // int + float64 연산 시 타입 에러
	// var f int = c // int 타입에 float64 대입 시도
	// b = 4
	// g := a * b // b는 int, a는 int64 연산 시 타입 에러

	/*
		타입 변환을 통해 해결 가능
	*/
	var e = float64(a) + c // int64를 float64로 변환하여 연산
	var f = a + int64(c)   // float64를 int64로 변환하여 연산 시 소수점 이하 절삭됨
	g := a * int64(b)      // int를 int64로 변환하여 연산

	fmt.Println(
		" vars e type:", reflect.TypeOf(e), "and value is:", e, "\n",
		"vars f type:", reflect.TypeOf(f), "and value is:", f, "\n",
		"vars g type:", reflect.TypeOf(g), "and value is:", g,
	)

	/*
		타입 변환 시 메모리 바이트 변경으로 값이 완전히 달라질 수 있음
	*/

	var h int16 = 3456
	var i int8 = int8(h) // int16를 int8로 변환 시 값이 달라짐

	fmt.Println(
		"vars h type:", reflect.TypeOf(h), "and value is:", h, "\n",
		"vars i type:", reflect.TypeOf(i), "and value is:", i,
	)
}
