package api

import (
	"electro_student/auth/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(userHandler *handlers.UsersHandler, roleHandler *handlers.RolesHandler, gradesHandler *handlers.GradesHandler) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/admin/add/user", userHandler.AddNewUser).Methods("POST")

	r.HandleFunc("/admin/add/role", roleHandler.AddNewRole).Methods("POST")

	r.HandleFunc("/admin/find/role/{id:[0-9]+}", roleHandler.FindRoleById).Methods("GET")

	r.HandleFunc("/admin/list/roles", roleHandler.ListRoles).Methods("GET")

	r.HandleFunc("/users/create", userHandler.Create).Methods("POST")

	r.HandleFunc("/users/find/{id:[0-9]+}", userHandler.Find).Methods("GET")

	r.HandleFunc("/users/find/{email}", userHandler.FindUserByEmail).Methods("GET")

	r.HandleFunc("/users/list", userHandler.ListUsers).Methods("GET")

	r.HandleFunc("/users/listby", userHandler.ListByParams).Methods("GET")

	r.HandleFunc("/users/update/pass", userHandler.UpdateUserPass).Methods("PUT")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()

	return r
}
