package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() == 1 { // 좌변이 falsy 이므로 이 조건은 절대 성립될 수 없음
		fmt.Println("1 increase")
	}
	if true || IncreaseAndReturn() == 2 { // 좌변이 truthy 이므로 <함수>는 절대 실행될 수 없음
		fmt.Println("2 increase")
	}

	fmt.Println("cnt:", cnt) // cnt 는 증가될 수 없음
}
