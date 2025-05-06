package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	// Float precision problem
	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c) // false
	fmt.Println(a + b)                                    // Float precision problem

	// Ignore the precision problem - bad way
	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, isEqual(a+b, c)) // true

	// Ignore the precision problem - good way
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, isEqualNextafter(a+b, c)) // true

	// Control precision manually like Financial calculation
	i, _ := new(big.Float).SetString("0.1")
	j, _ := new(big.Float).SetString("0.2")
	k, _ := new(big.Float).SetString("0.3")

	l := new(big.Float).Add(i, j) // 0.1 + 0.2
	fmt.Println(i, j, k, l)
	fmt.Println(k.Cmp(l)) // 0: 같음, 1: i가 더 큼, -1: j가 더 큼

}

const epsilon = 0.000001 // 무시할만한 상수 값

func isEqual(a, b float64) bool {
	/*
		Bad way, epsilon 값의 크기와 계산해야 하는 값의 크기를 고려해야함
	*/
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func isEqualNextafter(a, b float64) bool {
	/*
		Good way, 1bit 차이만큼 비교하여 같은 값이라고 간주
		a에서 b를 향해 1비트만 조정한 값 반환
	*/
	return math.Nextafter(a, b) == b
}
