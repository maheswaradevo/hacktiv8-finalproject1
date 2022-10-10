package main

import (
	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/router"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/server"
)

func main() {
	config.Init()
	cfg := config.GetConfig()
	root := mux.NewRouter()

	filename := "db/data.json"
	router.Init(root, filename)
	s := server.ProvideServer(cfg.ServerAddress, root)
	s.ListenAndServe()
}
