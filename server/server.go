package server

import (
	"fmt"
	"net/http"

	controllers "github.com/chris-weir/chrisweir/server/controllers/web"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func GetServer() *Server {
	server := Server{Router: chi.NewRouter()}

	return &server
}

func (server *Server) Run() error {
	server.MountHandlers()

	fmt.Println("Server starting on port :3000")
	return http.ListenAndServe(":3000", server.Router)
}

func (server *Server) MountHandlers() {
	server.mapWebRoutes()
	server.mapFileServer()
}

func (server *Server) mapWebRoutes() {
	server.Router.Get("/", controllers.Index)
	server.Router.NotFound(controllers.NotFound)
}

func (server *Server) mapFileServer() {
	assetsHandler := http.FileServer(http.Dir("assets"))
	server.Router.Get("/assets/*", http.StripPrefix("/assets", assetsHandler).ServeHTTP)
}
