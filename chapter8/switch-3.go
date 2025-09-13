package main

import "fmt"

func getMyAge() int {
	return 32
}

func main() {
	switch age := getMyAge(); true { // init 값 할당
	case age < 10:
		fmt.Println("아이")
	case age < 20:
		fmt.Println("학생")
	case age < 30:
		fmt.Println("20's")
	default:
		fmt.Println("조건에도 해당안되는 나이:", age)
	}
}
