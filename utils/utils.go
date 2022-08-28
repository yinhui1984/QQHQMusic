package utils

import "fmt"

var CurrentKeyWord string

const SearchCount = 50

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"

//var Yellow = "\033[33m"
//var Blue = "\033[34m"
//var Purple = "\033[35m"
//var Cyan = "\033[36m"
//var Gray = "\033[37m"
//var White = "\033[97m"

func GreenText(s string) string {
	return green + s + reset
}

func RedText(s string) string {
	return red + s + reset
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
