package utils

import "fmt"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintWelcome() {
	fmt.Println("Welcome~\n" +
		"Enter help for usage hints.")
}

func PrintAddPattern() {
	fmt.Println("Input your pattern, end with '$'.")
}

func PrintAddContent() {
	fmt.Println("Input your content, end with '$'.")
}

func PrintHelp() {
	fmt.Println("Usage:\n" +
		"  Add:\t<add a new interface> \n" +
		"  List:\t<list all interfaces>\n" +
		"  Delete:\t<delete one interface>")
}

func PrintAddDone() {
	fmt.Println("Add Done, the new interface has taken effect.")
}

func PrintDeleteDone() {
	fmt.Println("Delete Done, restart takes effect!")
}
