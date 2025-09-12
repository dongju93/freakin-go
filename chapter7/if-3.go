/*
[조건]
음식값이 오만 원이 넘고 친구중에 부자가 있다면 신발끈틀 묶는다. -> HasRichFriend()
부자가 없다면 돈을 나눠 낸다. -> HasRichFriend()
음식값이 3만원 이상 5만원 이하이고 같이 간 친구 수가 3명이 넘으면 신발끈을 묶는다. -> GetFriendCount()
3명 이하면 돈을 나눠 낸다. -> GetFriendCount()
3만원 미만이면 내가 낸다.
*/

package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendCount() int {
	return 3
}

func main() {

	price := 35000

	if price > 50000 && HasRichFriend() {
		fmt.Println("형님이시다, 신발끈을 묶어라")
	} else if price > 50000 {
		fmt.Println("1/N 하자")
	} else if price >= 30000 && GetFriendCount() > 3 {
		fmt.Println("눈치보고, 신발끈 묶어라")
	} else if price >= 30000 {
		fmt.Println("3인 이하 1/N 국룰")
	} else {
		fmt.Println("오늘은 내가 낸다")
	}

}
