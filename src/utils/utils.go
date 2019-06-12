package utils

import "fmt"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintWelcome() {
	fmt.Println("Welcome~")
}

func PrintAddPattern() {
	fmt.Println("Input your pattern.")
}

func PrintAddContent() {
	fmt.Println("Input your content.")
}

func PrintHelp() {
	fmt.Println("Usage:\n" +
		"  Add:\t<add a new interface> \n" +
		"  List:\t<list all interfaces>\n" +
		"  Delete:\t<delete one interface>")
}

func PrintAddDone() {
	fmt.Println("Add Done, restart takes effect！")
}

func PrintDeleteDone() {
	fmt.Println("Delete Done, restart takes effect!")
}
