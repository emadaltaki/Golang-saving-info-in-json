package main

import (
	"bufio"
	"fmt"
	"jsonfile/note/note"
	"os"
	"strings"
)

func main() {
	title, content := getNodeData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
}

func getNodeData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}
func getUserInput(prompt string) string {
	fmt.Println("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	// Checks only the last character of the string
	text = strings.TrimSuffix(text, "_world")

	//Short Cut to Trim the Spaces
	// text = strings.TrimSpace(text)
	return text
}
