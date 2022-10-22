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

var TodoList = &Todos{
	&Todo{
		ID:        1,
		Title:     "Develop Website",
		Details:   "Develop backend for company A",
		Completed: false,
		Priority:  1,
	},
}
