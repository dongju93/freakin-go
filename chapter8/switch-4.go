package main

import "fmt"

type ColorType int

const (
	RED ColorType = iota
	BLUE
	GREEN
	YELLOW
)

func colorToStringSwitch(color ColorType) string {
	switch color {
	case RED:
		return "Red"
	case BLUE:
		return "Blue"
	case GREEN:
		return "Green"
	case YELLOW:
		return "Yellow"
	default:
		return "Undefined"
	}

}

func colorToStringHashmap(color ColorType) string {
	colorMap := map[ColorType]string{ // switch 보다는 hashmap 이 더 적합해보여 추가로 구현해봄
		RED:    "Red",
		BLUE:   "Blue",
		GREEN:  "Green",
		YELLOW: "Yellow",
	}
	if strColor, isExist := colorMap[color]; isExist {
		return strColor
	}
	return "Undefined"
}

func getMyFavoriteColor() ColorType {
	return GREEN
}

func main() {
	fmt.Println("My favorite color is", colorToStringSwitch(getMyFavoriteColor()))
	fmt.Println("My favorite color is", colorToStringHashmap(getMyFavoriteColor()))
}
