package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
		break
	case 3:
		fmt.Println("a == 3")
		fallthrough // 다음 case 문까지 함께 실행, 코드 동작 파악에 혼란이 생길 수 있기때문에 사용 지양
	case 4:
		fmt.Println("a == 4")
	case 5:
		fmt.Println("a == 5")
	default:
		fmt.Println("a > 5")
	}
}
