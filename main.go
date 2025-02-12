package main

import (
	"bufio"
	"fmt"
	"jsonfile/note/note"
	"jsonfile/note/todo"
	"os"
	"strings"
)

// NOTE IF THE INTERFACE HAS ONLY ONE METHOD THEN U CALLED THE INTERFACE AS THE METHOD NAME + ER
type saver interface {
	Save() error
}

func main() {
	title, content := getNodeData()
	todoText := getUserInput("Todo text:")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed", err)
		return
	}
	fmt.Println("Todo saved successfully")
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed", err)
		return
	}
	fmt.Println("Note saved successfully")
}

func getNodeData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
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
