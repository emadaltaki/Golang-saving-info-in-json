package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)

}
func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)

	if err != nil {

		return err
	}
	fmt.Println("Note saved to json", json)
	return os.WriteFile(fileName, json, 0644)

}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("input cannot be empty")
	}

	return Todo{
		Text: content,
	}, nil

}
