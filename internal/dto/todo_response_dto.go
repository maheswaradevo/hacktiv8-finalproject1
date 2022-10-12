package dto

import "github.com/maheswaradevo/hacktiv8-finalproject1/internal/models"

type TodoResponse struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	Completed bool   `json:"completed"`
	Priority  int    `json:"priority"`
}
type TodoResponses []TodoResponse

func CreateTodoResponse(t models.Todo) TodoResponse {
	return TodoResponse{
		ID:        t.ID,
		Title:     t.Details,
		Details:   t.Details,
		Completed: t.Completed,
		Priority:  t.Priority,
	}
}

func CreateTodoResponses(t models.Todos) *TodoResponses {
	var todoResponses TodoResponses

	for _, idx := range t {
		todo := CreateTodoResponse(*idx)
		todoResponses = append(todoResponses, todo)
	}
	return &todoResponses
}

func CreateTodoByIDResponses(t models.Todo) *TodoResponse {
	return &TodoResponse{
		ID:        t.ID,
		Title:     t.Title,
		Details:   t.Details,
		Completed: t.Completed,
		Priority:  t.Priority,
	}
}

func UpdateTodoResponses(t models.Todo) *TodoResponse {
	return &TodoResponse{
		ID:        t.ID,
		Title:     t.Title,
		Details:   t.Details,
		Completed: t.Completed,
		Priority:  t.Priority,
	}
}

func DeleteTodoResponses(t models.Todo) TodoResponse {
	return TodoResponse{
		ID:        t.ID,
		Title:     t.Title,
		Details:   t.Details,
		Completed: t.Completed,
		Priority:  t.Priority,
	}
}

func CreateDeleteTodoResponses(t models.Todos) *TodoResponses {
	var todoResponses TodoResponses

	for _, idx := range t {
		todo := DeleteTodoResponses(*idx)
		todoResponses = append(todoResponses, todo)
	}
	return &todoResponses
}
