package model

import "fmt"

type Todo struct {
	Id   int    `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func (todo *Todo) String() string {
	return fmt.Sprintf("{id: %v, todo: %v, done: %v}", todo.Id, todo.Todo, todo.Done)
}
