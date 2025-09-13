package main

import "fmt"

func main() {
	temp := 12

	switch true { // switch 각 문을 조건문으로 사용, true 인 경우를 찾음
	case temp < 10, temp > 30:
		fmt.Println("야외 활동하기 좋은 날씨가 아니에요")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 겉옷을 챙기세요")
	case temp >= 15 && temp < 25:
		fmt.Println("야외 활동하기 딱 좋은 날씨에요")
	default:
		fmt.Println("따뜻한 날씨 입니다")
	}
}
