package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	routes.HandleFunc("/{id}", t.getTodoByID()).Methods(http.MethodGet)
	routes.HandleFunc("/{id}", t.updateTodo()).Methods(http.MethodPut)
	routes.HandleFunc("/{id}", t.deleteTodo()).Methods(http.MethodDelete)
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

func (t *todoHandler) getTodoByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryVar := mux.Vars(r)
		id := queryVar["id"]
		idConv, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Printf("[getTodoByID] failed to convert id to uint: %v", err)
			utils.NewErrorResponse(w, err)
			return
		}
		res, err := t.ts.GetTodoByID(r.Context(), idConv)
		if err != nil {
			log.Printf("[getTodoByID] failed to get todo by id, id : %v, err : %v", idConv, err)
			utils.NewErrorResponse(w, err)
			return
		}
		utils.NewSuccessResponsWriter(w, http.StatusOK, "SUCCESS", res)
	}
}

func (t *todoHandler) updateTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := dto.TodoRequest{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Printf("[updateTodo] failed to parse JSON data: %v", err)
			utils.NewErrorResponse(w, err)
			return
		}
		queryVar := mux.Vars(r)
		id := queryVar["id"]
		idConv, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Printf("[getTodoByID] failed to convert id to uint: %v", err)
			utils.NewErrorResponse(w, err)
			return
		}
		res, err := t.ts.UpdateData(r.Context(), idConv, &data)
		if err != nil {
			log.Printf("[getTodoByID] failed to get todo by id, id : %v, err : %v", idConv, err)
			utils.NewErrorResponse(w, err)
			return
		}
		utils.NewSuccessResponsWriter(w, http.StatusOK, "SUCCESS", res)
	}
}

func (t *todoHandler) deleteTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryVar := mux.Vars(r)
		id := queryVar["id"]
		idConv, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Printf("[deleteTodoByID] failed to convert id to uint: %v", err)
			utils.NewErrorResponse(w, err)
			return
		}
		res, err := t.ts.DeleteData(r.Context(), idConv)
		if err != nil {
			log.Printf("[deleteTodoByID] failed to delete the order by id, err => %v, id => %v", err, idConv)
			utils.NewErrorResponse(w, err)
			return
		}
		utils.NewSuccessResponsWriter(w, http.StatusOK, "SUCCESS", res)
	}
}
