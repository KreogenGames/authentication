package api

import (
	"electro_student/auth/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(userHandler *handlers.UsersHandler) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/users/create", userHandler.Create).Methods("POST")

	r.HandleFunc("/users/list", userHandler.List).Methods("GET")

	r.HandleFunc("/users/find/{id:[0-9]+}", userHandler.Find).Methods("GET")

	r.HandleFunc("/users/find/email/{email:[A-Z]@[a-z].ru}", userHandler.FindUserByEmail).Methods("GET")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()

	return r
}
