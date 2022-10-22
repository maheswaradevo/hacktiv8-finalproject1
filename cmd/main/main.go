package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/maheswaradevo/hacktiv8-finalproject1/docs"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/router"
	"github.com/maheswaradevo/hacktiv8-finalproject1/internal/global/server"
	"github.com/maheswaradevo/hacktiv8-finalproject1/pkg/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initializeGlobalRouter(whitelisted string) *mux.Router {
	r := mux.NewRouter()

	arrayWhiteListedUrls := strings.Split(whitelisted, ",")
	for idx := range arrayWhiteListedUrls {
		log.Printf(arrayWhiteListedUrls[idx])
	}
	whiteListedUrls := make(map[string]bool)

	for _, v := range arrayWhiteListedUrls {
		whiteListedUrls[v] = true
	}
	r.Use(middleware.CorsMiddleware(whiteListedUrls))
	return r
}

// @title Todo API
// @version 1.0
// @description This is an API to managing our Todo List
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host https://todos-api-go.herokuapp.com
// @BasePath /
func main() {
	config.Init()
	cfg := config.GetConfig()
	root := initializeGlobalRouter(cfg.WhiteListed)
	filename := "db/data.json"

	router.Init(root, filename)
	root.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	port := fmt.Sprintf("%s:%s", "0.0.0.0", cfg.PORT)
	s := server.ProvideServer(port, root)
	s.ListenAndServe()
}
