package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	if age, ok := getMyAge(); ok && age < 20 { // age, ok 초기 값 할당
		fmt.Println("학생입니다.", age)
	} else if ok && age < 30 {
		fmt.Println("청년입니다.", age)
	} else if ok {
		fmt.Println("중장년입니다.", age)
	} else {
		fmt.Println("나이를 알 수 없습니다.")
	}

	// fmt.Println("당신의 나이는 ", age) // age 변수는 if 블록 안에서만 유효함
}
