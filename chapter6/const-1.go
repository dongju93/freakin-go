package chapter6

const (
	Red   int = iota // 0부터 시작해 1씩 증가
	Blue  int = iota
	Green int = iota
) // 소괄호를 벗어나면 iota 초기화

const (
	C1 uint = iota + 1
	C2      // 같은 규칙이 적용된다면 타입과 iota 생략 가능
	C3
)

const (
	BitFlag1 uint = 1 << iota // iota=0, 1<<0 = 1 (이진수: 0001, 첫 번째 비트가 1)
	// `<<` 는 왼쪽 시프트 연산자로, 1을 iota만큼 왼쪽으로 이동시키고 오른쪽에 0을 채웁니다.
	BitFlag2 // iota=1, 1<<1 = 2 (이진수: 0010, 두 번째 비트가 1)
	BitFlag3 // iota=2, 1<<2 = 4 (이진수: 0100, 세 번째 비트가 1)
	BitFlag4 // iota=3, 1<<3 = 8 (이진수: 1000, 네 번째 비트가 1)
)
