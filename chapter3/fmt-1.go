package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var c int = 43431
	var d float64 = 324.13455
	var e float64 = 3.14
	var f float64 = 32799438743.8287

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f)
	fmt.Printf("a: %d b: %d f: %f\n\n", a, b, f)
	fmt.Println("=====서식 지정자=====")
	fmt.Printf("데이터 타입 기본: %v\n", f)
	fmt.Printf("데이터 타입: %T\n", f)
	fmt.Printf("Boolean: %t\n", true)
	fmt.Printf("10진수: %d\n", 100)
	fmt.Printf("2진수: %b\n", 100)
	fmt.Printf("8진수: %o\n", 100)
	fmt.Printf("8진수임을 표시: %#o\n", 100)
	fmt.Printf("16진수: %x\n", 12344321)
	fmt.Printf("16진수 대문자: %X\n", 12344321)
	fmt.Printf("유니코드: %c\n", 100)
	fmt.Printf("지수 표현: %e\n", f)
	fmt.Printf("지수가 아닌 표현: %f\n", f)
	fmt.Printf("값이 크면 지수, 작으면 실수 표현: %g\n", f)
	fmt.Printf("문자열: %s\n", "Hello, world!")
	fmt.Printf("문자열 특수 기호 제외: %q\n", "Hello, \nworld!\n")
	fmt.Printf("포인터: %p\n\n", &a)
	fmt.Println("=====출력 너비 지정=====")
	fmt.Printf("최소 너비 지정: %5d, %5d\n", a, b)
	fmt.Printf("최소 너비 지정 0채우기: %05d, %05d\n", a, b)
	fmt.Printf("최소 너비 지정 왼쪽 정렬: %-5d, %-5d\n", a, b)
	fmt.Printf("최소 너비 지정: %5d\n", c)
	fmt.Printf("최소 너비 지정 0채우기: %05d\n", c)
	fmt.Printf("최소 너비 지정 왼쪽 정렬: %-5d\n\n", c)
	fmt.Printf("최소 너비 지정 0채우기, 소수점 이하 2자리: %08.2f\n", d)
	fmt.Printf("최소 너비 지정 0채우기, 총 숫자 2자리: %08.2g\n", d)
	fmt.Printf("최소 너비 지정, 총 숫자 5자리: %8.5g\n", d)
	fmt.Printf("소수점 이하 6자리: %f\n", e)
	fmt.Println("=====특수문자 출력=====")
	stringify_special_char := "Hello\tGo\tWorld\n\"Go\"is Awesome!\n"
	fmt.Print(stringify_special_char)
	fmt.Printf("%s", stringify_special_char)
	fmt.Printf("%q", stringify_special_char)
}
