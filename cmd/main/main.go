package main

import (
	"github.com/gorilla/mux"
	_ "github.com/maheswaradevo/hacktiv8-finalproject1/docs"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/router"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/server"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Todo API
// @version 1.0
// @description This is an API to managing our Todo List
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
	config.Init()
	cfg := config.GetConfig()
	root := mux.NewRouter()

	filename := "db/data.json"
	router.Init(root, filename)
	root.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	s := server.ProvideServer(cfg.ServerAddress, root)
	s.ListenAndServe()
}
