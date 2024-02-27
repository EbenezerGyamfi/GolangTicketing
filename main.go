package main

import (
	"errors"
	"fmt"
)

func main()  {
	
	
	title, content, errr := getUserResult()

	if(errr != nil){
		fmt.Println("Error occured")
	}

	fmt.Println(title, content)


}

func getUserResult() (string, string, error)  {
	title,  err := getUserData("Enter your title")

	if(err != nil){
		panic("Error occured")
	}

	content, err := getUserData("Enter your content")

	if(err != nil){
		fmt.Println("Error occured")
	}

	return title, content, nil
}

func getUserData(plainMessage string) (string, error)  {
	var userInput string
	fmt.Println(plainMessage)

	 fmt.Scanln(&userInput)

	 if(userInput == ""){
		return "", errors.New(" empty title or content")
	 }

	return userInput, nil
}