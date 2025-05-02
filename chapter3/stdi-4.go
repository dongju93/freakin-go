package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin) // 1. 인수로 입력되는 input stream 으로 Reader 객체 생성

	var a int
	var b int

	println("정수 두 개를 입력하세요.")
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n') // 2. 현재 줄의 끝(줄바꿈 문자)까지 모든 문자를 읽어들여, 입력 버퍼에서 제거
	} else {
		fmt.Println(n, a, b)
	}

	println("정수 두 개를 다시 입력하세요.")
	n, err = fmt.Scanln(&a, &b) // 3. 다시 입력받기
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

}
