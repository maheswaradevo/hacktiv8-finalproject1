package todo

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/constant"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/utils"
)

type todoHandler struct {
	r  *mux.Router
	ts TodoService
}

func (t *todoHandler) InitHandler() {
	routes := t.r.PathPrefix(constant.TODO_API_PATH).Subrouter()

	routes.HandleFunc("", t.getAllTodos()).Methods(http.MethodGet)
	routes.HandleFunc("", t.createTodo()).Methods(http.MethodPost)
}

func ProvideTodoHandler(r *mux.Router, ts TodoService) *todoHandler {
	return &todoHandler{
		r:  r,
		ts: ts,
	}
}

func (t *todoHandler) getAllTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := t.ts.GetAllData(r.Context())
		if err != nil {
			utils.NewErrorResponse(w, err)
			return
		}
		utils.NewSuccessResponsWriter(w, http.StatusOK, "SUCCESS", res)
	}
}

func (t *todoHandler) createTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := dto.TodoRequest{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Printf("[createTodo] failed to parse JSON data: %v", err)
			utils.NewErrorResponse(w, err)
			return
		}
		t.ts.CreateNewData(r.Context(), &data)
		utils.NewSuccessResponsWriter(w, http.StatusCreated, "SUCCESS", data.ID)
	}
}
