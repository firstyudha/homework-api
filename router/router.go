package router

import (
	"net/http"
	"os"

	"github.com/rysmaadit/go-template/handler"
	"github.com/rysmaadit/go-template/service"
	"gorm.io/gorm"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(dependencies service.Dependencies, db *gorm.DB) http.Handler {
	r := mux.NewRouter()

	setAuthRouter(r, dependencies.AuthService)
	// pass object db ke turunannya
	setMovieRouter(r, db)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

// setup router untuk movies
func setMovieRouter(router *mux.Router, db *gorm.DB) {
	// pass object db ke handler
	router.Methods(http.MethodPost).Path("/movies").Handler(handler.CreateMovie(db))
	router.Methods(http.MethodGet).Path("/movies/{slug}").Handler(handler.GetMovie(db))
	router.Methods(http.MethodGet).Path("/movies").Handler(handler.GetMovies(db))
	router.Methods(http.MethodPut).Path("/movies/{slug}").Handler(handler.UpdateMovie(db))
	router.Methods(http.MethodDelete).Path("/movies/{slug}").Handler(handler.DeleteMovie(db))
}

func setAuthRouter(router *mux.Router, dependencies service.AuthServiceInterface) {
	router.Methods(http.MethodGet).Path("/auth/token").Handler(handler.GetToken(dependencies))
	router.Methods(http.MethodPost).Path("/auth/token/validate").Handler(handler.ValidateToken(dependencies))
}
