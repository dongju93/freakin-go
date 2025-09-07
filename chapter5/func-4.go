package main

import "fmt"

func printNo(n int) {
	if n == 0 {
		return // 종료조건이 없으면 무한 반복되어 비정상 종료됨
	}
	fmt.Println(n)
	printNo(n - 1)          // Recursive
	fmt.Println("After", n) // Stack(FILO) 으로 n 출력, 0은 이미 return 되었기때문에 미출력
}

func main() {
	printNo(3)
}

/*
output:
3
2
1
After 1
After 2
After 3
*/
