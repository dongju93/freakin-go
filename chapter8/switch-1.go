package main

import "fmt"

func main() {
	day := "thursday"

	switch day {
	case "monday", "tuesday": // 여러 값을 switch 조건에 포함시킬 경우 , 로 구분
		fmt.Println("월요일, 화요일은 짜파게티 요리사")
	case "wednesday", "thursday":
		fmt.Println("수요일, 목요일은 간짬뽕 요리사")
	case "friday":
		fmt.Println("금요일은 불닭볶음면 요리사")
	}
}
