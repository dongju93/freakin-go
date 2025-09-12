package main

import "fmt"

func main() {
	var age = 22

	if age >= 10 && age <= 15 { // and 조건: 10 이상이면서 15이하
		fmt.Println("You are young")
	} else if age > 30 || age < 20 { // or 조건: 30 초과거나 20 미만
		fmt.Println("You are not 20s")
	} else {
		fmt.Println("Best age of your life")
	}
}
