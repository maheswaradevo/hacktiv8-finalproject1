package router

import (
	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/todo"
)

func Init(router *mux.Router, filename string) {
	webRouter := router.NewRoute().PathPrefix("/api/v1").Subrouter()

	todoService := todo.ProvideTodoService(filename)
	todoHandler := todo.ProvideTodoHandler(webRouter, todoService)
	todoHandler.InitHandler()
}
