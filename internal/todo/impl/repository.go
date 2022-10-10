package impl

import (
	"context"
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/models"
)

type TodoRepository interface {
	GetAllData(ctx context.Context) (models.Todos, error)
	CreateNewData(ctx context.Context, reqData *dto.TodoRequest) error
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
	dataByte, err := json.Marshal(models.TodoList)
	if err != nil {
		log.Printf("[GetAllData] cannot marshalling the struct: %v", err)
		return nil, err
	}
	err = ioutil.WriteFile(t.filename, dataByte, fs.ModeAppend)
	if err != nil {
		log.Printf("[GetAllData] cannot write to JSON: %v", err)
		return nil, err
	}
	return models.TodoList, nil
}

func (t TodoRepositoryImpl) CreateNewData(ctx context.Context, reqData *dto.TodoRequest) error {
	dataByte, err := json.Marshal(reqData)
	if err != nil {
		log.Printf("[CreateNewData] cannot marshalling the struct: %v", err)
		return err
	}
	err = ioutil.WriteFile(t.filename, dataByte, fs.ModeAppend)
	if err != nil {
		log.Printf("[CreateNewData] cannot write to JSON: %v", err)
		return err
	}
	models.TodoList = append(models.TodoList, (*models.Todo)(reqData))
	return nil
}
