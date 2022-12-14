package dto

import "github.com/maheswaradevo/hacktiv8-finalproject1/internal/models"

type TodoRequest struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	Completed bool   `json:"completed"`
	Priority  int    `json:"priority"`
}

func (tr *TodoRequest) ToEntity() (t *models.Todo) {
	t = &models.Todo{
		ID:        tr.ID,
		Title:     tr.Title,
		Details:   tr.Details,
		Completed: tr.Completed,
		Priority:  tr.Priority,
	}
	return
}
