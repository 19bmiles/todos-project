package graph

import (
	"github.com/19bmiles/todos-project/pkg/todo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	ToDo todo.Todo
}
