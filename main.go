package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNodeData()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNodeData() (string, string, error) {
	title, err := getUserInput("Note title:")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	content, err := getUserInput("Note content:")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return title, content, nil
}
func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)

	if input == "" {
		return "", errors.New("Input cannot be empty")
	}
	return input, nil
}
