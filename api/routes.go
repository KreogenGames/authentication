package api

import (
	"electro_student/auth/internals/app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoutes(userHandler *handlers.UsersHandler, roleHandler *handlers.RolesHandler, gradesHandler *handlers.GradesHandler) *mux.Router {

	r := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./internals/ui/static/"))

	r.HandleFunc("/", handlers.HomePage)

	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", fileServer))

	r.HandleFunc("/admin", handlers.AdminPage)

	r.HandleFunc("/admin/roles", handlers.RoleStatPage)

	r.HandleFunc("/admin/users", handlers.UserStatPage)

	r.HandleFunc("/admin/add/user", userHandler.AddNewUser).Methods("POST") //На страницу админа

	r.HandleFunc("/admin/add/role", roleHandler.AddNewRole).Methods("POST") //На страницу админа

	r.HandleFunc("/admin/find/role/{id:[0-9]+}", roleHandler.FindRoleById).Methods("GET")

	r.HandleFunc("/admin/list/roles", roleHandler.ListRoles).Methods("GET") //На страницу админа

	r.HandleFunc("/users/create", userHandler.Create).Methods("POST") //На страницу регистрации

	r.HandleFunc("/users/find/{id:[0-9]+}", userHandler.Find).Methods("GET") //На страницу пользователя

	r.HandleFunc("/users/find/{email}", userHandler.FindUserByEmail).Methods("GET") //На страницу админа

	r.HandleFunc("/users/list", userHandler.ListUsers).Methods("GET") //На страницу админа

	r.HandleFunc("/users/listby", userHandler.ListByParams).Methods("GET") //На страницу пользователя, как поисковик

	r.HandleFunc("/users/update/pass", userHandler.UpdateUserPass).Methods("PUT") //На страницу обновления пароля

	r.HandleFunc("/disciplines/disciplineid/grades", gradesHandler.Create).Methods("POST") //На страницу преподавателя

	r.HandleFunc("/disciplines/list/grades", gradesHandler.List).Methods("GET") //На страницу админа, студента и препода

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()

	return r
}
