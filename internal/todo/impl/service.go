package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/models"
)

type TodoServiceImpl struct {
	repo TodoRepository
}

func ProvideOrderService(repo TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{
		repo: repo,
	}
}

func (t TodoServiceImpl) GetAllData(ctx context.Context) (*dto.TodoResponses, error) {
	check, err := t.repo.CheckTodo(ctx)
	if !check {
		log.Printf("[GetAllData] there's no data")
		return nil, err
	}
	res, err := t.repo.GetAllData(ctx)
	if err != nil {
		log.Printf("[GetAllData] failed to get all the data, err => %v", err)
		return nil, err
	}
	return dto.CreateTodoResponses(res), nil
}

func (t TodoServiceImpl) CreateNewData(ctx context.Context, data *dto.TodoRequest) error {
	for idx := 0; idx < len(models.TodoList); idx++ {
		if data.ID == models.TodoList[idx].ID {
			data.ID = data.ID + 1
		}
	}
	err := t.repo.CreateNewData(ctx, data)
	if err != nil {
		log.Printf("[CreateNewData] failed to create a new data: %v", err)
		return err
	}
	return nil
}

func (t TodoServiceImpl) GetTodoByID(ctx context.Context, id uint64) (*dto.TodoResponse, error) {
	check, err := t.repo.CheckTodoByID(ctx, id)
	if !check {
		log.Printf("[GetTodoByID] there's no todo with id: %v", id)
		return nil, err
	}
	res, err := t.repo.GetTodoByID(ctx, id)
	if err != nil {
		log.Printf("[GetTodoByID] failed to get todo by id, id => %v, err => %v", id, err)
		return nil, err
	}

	return dto.CreateTodoByIDResponses(res), nil
}
