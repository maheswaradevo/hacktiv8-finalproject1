package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/dto"
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
	res, err := t.repo.GetAllData(ctx)
	if err != nil {
		log.Printf("[GetAllData] failed to get all the data, err => %v", err)
		return nil, err
	}
	return dto.CreateTodoResponses(res), nil
}

func (t TodoServiceImpl) CreateNewData(ctx context.Context, data *dto.TodoRequest) error {
	err := t.repo.CreateNewData(ctx, data)
	if err != nil {
		log.Printf("[CreateNewData] failed to create a new data: %v", err)
		return err
	}
	return nil
}
