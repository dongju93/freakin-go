package main

import "fmt"

func main() {
	fmt.Println("======4.1.1 연산의 결과 타입======")
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y =", x+y)
	fmt.Println("x - y =", x-y)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / y =", x/y)
	fmt.Println("x % y =", x%y)

	fmt.Println("s * t =", s*t)
	fmt.Println("s / t =", s/t)

	fmt.Println("======4.1.2 비트 연산자======")

	var x1 int8 = 34   // 8비트 정수, 00100010
	var x2 int16 = 34  // 16비트 정수, 00000000 00100010
	var x3 uint8 = 34  // 8비트 부호 없는 정수, 00100010
	var x4 uint16 = 34 // 16비트 부호 없는 정수, 00000000 00100010
	var mask int8 = 2  // 00000010

	fmt.Printf("^%d = %5d,\t %08b\n", x1, ^x1, uint8(^x1))   // 2진수 출력, 비트 반전 호 출력, 비트 반전 후 이진수 출력
	fmt.Printf("^%d = %5d,\t %016b\n", x2, ^x2, uint16(^x2)) // 메모리 비트 패턴을 그대로 확인하려면 부호 없는 정수 타입(uint)으로 형변환
	fmt.Printf("^%d = %5d,\t %08b\n", x3, ^x3, uint8(^x3))
	fmt.Printf("^%d = %5d,\t %016b\n", x4, ^x4, ^x4)
	// 비트 클리어
	/*
		34 (십진수):  0 0 1 0 0 0 1 0 (이진수)
		2 (십진수):   0 0 0 0 0 0 1 0 (이진수)
		              -----------------
		34&^2 결과:   0 0 1 0 0 0 0 0 = 32 (십진수)
		                         ↑
		                특정 비트가 0으로 클리어됨
	*/
	fmt.Printf("%d&^%d = %5d,\t %08b\n", x1, mask, x1&^mask, uint8(x1&^mask)) // 특정 비트만 0으로 설정

	fmt.Println("======4.1.3 시프트 연산자======")
	// 왼쪽 혹은 오른쪽으로 밀거나 당기는 연산자
	var i int8 = 4  // 8비트 정수, 00000100
	var j int8 = 64 // 8비트 정수, 01000000

	fmt.Printf("x:%08b x<<2: %08b x<<2: %d\n", i, i<<2, i<<2) // 왼쪽으로 2비트 Shift
	fmt.Printf("y:%08b y<<2: %08b y<<2: %d\n", j, j<<2, j<<2) // 값이 버려지지 않은지 확인해야함
	fmt.Printf("x:%08b x>>2: %08b x>>2: %d\n", i, i>>2, i>>2) // 오른쪽으로 2비트 Shift
	fmt.Printf("y:%08b y>>2: %08b y>>2: %d\n", j, j>>2, j>>2)

	// 음수, 양수에 따른 오른쪽 시프트
	var k int8 = 16
	var g int8 = -128
	var z int8 = -1
	var w uint8 = 128
	fmt.Printf("k:%08b k>>2: %08b k>>2: %d\n", k, k>>2, k>>2) // 오른쪽으로 2비트 Shift
	fmt.Printf("g:%08b g>>2: %08b g>>2: %d\n", uint8(g), uint8(g>>2), g>>2)
	fmt.Printf("z:%08b z>>2: %08b z>>2: %d\n", uint8(z), uint8(z>>2), z>>2)
	fmt.Printf("w:%08b w>>2: %08b w>>2: %d\n", uint8(w), uint8(w>>2), w>>2)

}
