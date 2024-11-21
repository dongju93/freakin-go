/*
이 코드가 어떤 패키지에 속해 있는지 선언
go 언어의 모든 코드는 패키지 선언으로 시작해야 함

main() 함수가 없는 패키지는 패키지 이름으로 main 을 쓸 수 없음
곧, 실행 파일을 만들 수 없고 다른 패키지에서 import 할 함수로 사용됨
*/
package main

// 표준 입출력 내장 패키지
import "fmt"

/*
프로그램의 진입점
main() 함수는 프로그램의 시작과 끝을 나타냄
*/
func main() {
	fmt.Println("Hello, World!")
}
