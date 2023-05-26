package handlers

import (
	"electro_student/auth/internals/app/models"
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type templateData struct {
	Roles []*models.Role
	Users []*models.User
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./internals/ui/html/home.page.tmpl",
		"./internals/ui/html/base.layout.tmpl",
		"./internals/ui/html/header.partial.tmpl",
		"./internals/ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func AdminPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./internals/ui/html/admin.panel.tmpl",
		"./internals/ui/html/header.admin.partial.tmpl",
		"./internals/ui/html/nav.admin.partial.tmpl",
		"./internals/ui/html/footer.partial.tmpl",
		"./internals/ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (handler *RolesHandler) RoleStatPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/roles" {
		http.NotFound(w, r)
		return
	}

	list := handler.processor.SliceRoles()

	data := &templateData{Roles: list}

	files := []string{
		"./internals/ui/html/admin.panel.tmpl",
		"./internals/ui/html/role.list.tmpl",
		"./internals/ui/html/header.admin.partial.tmpl",
		"./internals/ui/html/nav.admin.partial.tmpl",
		"./internals/ui/html/footer.partial.tmpl",
		"./internals/ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (handler *UsersHandler) UserStatPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/users" {
		http.NotFound(w, r)
		return
	}

	list := handler.processor.SliceUsers()

	data := &templateData{Users: list}

	files := []string{
		"./internals/ui/html/admin.panel.tmpl",
		"./internals/ui/html/users.list.tmpl",
		"./internals/ui/html/header.admin.partial.tmpl",
		"./internals/ui/html/nav.admin.partial.tmpl",
		"./internals/ui/html/footer.partial.tmpl",
		"./internals/ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
