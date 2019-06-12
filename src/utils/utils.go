package utils

import "fmt"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintTips() {
	fmt.Println("Welcome~\n" +
		"input your pattern.")
}
