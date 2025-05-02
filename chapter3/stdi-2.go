package main

import "fmt"

func main() {
	var a int // 1. 값을 저장할 변수
	var b int

	println("정수 두 개를 입력하세요.")
	n, err := fmt.Scanf("%d %d\n", &a, &b) // 2. 한 줄에 띄어쓰기로 두 개를 구분
	if err != nil {                        // 3. 에러가 발생하면 에러 코드 출력
		fmt.Println(n, err)
	} else { // 4. 정상 입력되면 입력 값 출력
		fmt.Println(n, a, b)
	}

}
