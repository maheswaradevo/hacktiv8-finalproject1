package models

// Todo represent the model for a Todo
type Todo struct {
	ID        uint64
	Title     string
	Details   string
	Completed bool
	Priority  int
}

type Todos []*Todo

var TodoList = &Todos{}
