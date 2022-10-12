package todo

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/todo/impl"
)

type TodoService interface {
	GetAllData(ctx context.Context) (*dto.TodoResponses, error)
	CreateNewData(ctx context.Context, data *dto.TodoRequest) error
	GetTodoByID(ctx context.Context, id uint64) (*dto.TodoResponse, error)
	UpdateData(ctx context.Context, id uint64, data *dto.TodoRequest) (*dto.TodoResponse, error)
	DeleteData(ctx context.Context, id uint64) (*dto.TodoResponses, error)
}

func ProvideTodoService(filename string) TodoService {
	repo := impl.ProvideTodoRepository(filename)
	return impl.ProvideOrderService(repo)
}
