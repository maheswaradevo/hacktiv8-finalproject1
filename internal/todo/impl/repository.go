package impl

import (
	"context"
	"errors"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/models"
)

type TodoRepository interface {
	GetAllData(ctx context.Context) (models.Todos, error)
	CreateNewData(ctx context.Context, reqData *dto.TodoRequest) error
	GetTodoByID(ctx context.Context, id uint64) (models.Todo, error)
	UpdateData(ctx context.Context, id uint64, reqData *dto.TodoRequest) (models.Todo, error)
	CheckTodoByID(ctx context.Context, id uint64) (bool, error)
	CheckTodo(ctx context.Context) (bool, error)
	DeleteData(ctx context.Context, id uint64) (models.Todos, error)
}

type TodoRepositoryImpl struct {
	filename string
}

func ProvideTodoRepository(filename string) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		filename: filename,
	}
}

func (t TodoRepositoryImpl) GetAllData(ctx context.Context) (models.Todos, error) {
	err := utils.WriteToJSON(models.TodoList, t.filename)
	if err != nil {
		log.Printf("[GetAllData] failed writing to JSON: %v", err)
		return nil, err
	}
	return *models.TodoList, nil
}

func (t TodoRepositoryImpl) CreateNewData(ctx context.Context, reqData *dto.TodoRequest) error {
	todos := *models.TodoList

	todos = append(todos, (*models.Todo)(reqData))
	err := utils.WriteToJSON(models.TodoList, t.filename)
	if err != nil {
		log.Printf("[CreateNewData] failed to writing to JSON: %v", err)
		return err
	}
	models.TodoList = &todos
	return nil
}

func (t TodoRepositoryImpl) GetTodoByID(ctx context.Context, id uint64) (models.Todo, error) {
	var todoByID models.Todo
	for _, todo := range *models.TodoList {
		if todo.ID == id {
			todoByID = *todo
		}
	}
	err := utils.WriteToJSON(todoByID, t.filename)
	if err != nil {
		log.Printf("[GetTodoByID] failed writing to JSON: %v", err)
		return models.Todo{}, err
	}
	return todoByID, nil
}

func (t TodoRepositoryImpl) CheckTodoByID(ctx context.Context, id uint64) (bool, error) {
	for _, todo := range *models.TodoList {
		if todo.ID == id {
			return true, nil
		}
	}
	errDataNotExists := errors.New("data doesn't exists")
	return false, errDataNotExists
}

func (t TodoRepositoryImpl) CheckTodo(ctx context.Context) (bool, error) {
	for _, todo := range *models.TodoList {
		if todo.ID != 0 {
			return true, nil
		}
	}
	errDataNotExists := errors.New("data doesn't exists")
	return false, errDataNotExists
}

func (t TodoRepositoryImpl) UpdateData(ctx context.Context, id uint64, reqData *dto.TodoRequest) (models.Todo, error) {
	var todoByID models.Todo
	for _, todo := range *models.TodoList {
		if todo.ID == id {
			todoByID = *reqData.ToEntity()
			todoByID.ID = id
			todo.Title = reqData.Title
			todo.Details = reqData.Details
			todo.Completed = reqData.Completed
			todo.Priority = reqData.Priority
		}
	}
	err := utils.WriteToJSON(models.TodoList, t.filename)
	if err != nil {
		log.Printf("[UpdateData] failed to writing to JSON: %v", err)
		return models.Todo{}, err
	}
	return todoByID, nil
}

func (t TodoRepositoryImpl) DeleteData(ctx context.Context, id uint64) (models.Todos, error) {
	todos := *models.TodoList

	idx := id - 1

	todosModify := append(todos[:idx], todos[idx+1:]...)

	err := utils.WriteToJSON(models.TodoList, t.filename)
	if err != nil {
		log.Printf("[DeleteData] failed to writing to JSON: %v", err)
	}
	models.TodoList = &todosModify
	return todosModify, nil
}
