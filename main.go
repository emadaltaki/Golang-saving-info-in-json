package main

import (
	"fmt"
	"jsonfile/note/note"
)

func main() {
	title, content := getNodeData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNodeData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)

	return input
}
