package main

import (
	"fmt"
	"github.com/ebenezer/note"
)

func main() {

	title, content := getUserResult()

	note, errr := note.New(title, content)

	if errr != nil {
		fmt.Println("Error occured")
	}

	fmt.Println(note)

}

func getUserResult() (string, string) {
	title := getUserData("Enter your title")

	content := getUserData("Enter your content")

	return title, content
}

func getUserData(plainMessage string) string {
	var userInput string
	fmt.Println(plainMessage)

	fmt.Scanln(&userInput)

	return userInput
}
