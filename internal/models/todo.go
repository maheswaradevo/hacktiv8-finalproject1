package models

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
		Title:     "Test1",
		Details:   "Test1",
		Completed: false,
		Priority:  1,
	},
	&Todo{
		ID:        2,
		Title:     "Test2",
		Details:   "Test2",
		Completed: false,
		Priority:  1,
	},
	&Todo{
		ID:        3,
		Title:     "Test3",
		Details:   "Test3",
		Completed: false,
		Priority:  2,
	},
	&Todo{
		ID:        4,
		Title:     "Test4",
		Details:   "Test4",
		Completed: false,
		Priority:  2,
	},
}
