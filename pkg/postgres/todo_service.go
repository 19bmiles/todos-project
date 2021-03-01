package postgres

import "fmt"

type ToDoService struct {
	Message string
}

func (t *ToDoService) Create(text string, isDone bool)  {
	fmt.Sprintf("I have a message: %s\n", t.Message)
	panic("oops")
}