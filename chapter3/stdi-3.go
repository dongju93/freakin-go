package main

import "fmt"

func main() {
	var a int // 1. 값을 저장할 변수
	var b int

	println("정수 두 개를 입력하세요.")
	n, err := fmt.Scanln(&a, &b) // 2. 입력 두 개를 한줄에 받기, 공백으로 구분, 엔터 키를 누르면 입력 종료
	if err != nil {              // 3. 에러가 발생하면 에러 코드 출력
		fmt.Println(n, err)
	} else { // 4. 정상 입력되면 입력 값 출력
		fmt.Println(n, a, b)
	}

}
