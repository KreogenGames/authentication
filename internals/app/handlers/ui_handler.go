package handlers

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./internals/ui/html/home.page.tmpl",
		"./internals/ui/html/base.layout.tmpl",
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

func RoleStatPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/roles" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./internals/ui/html/admin.panel.tmpl",
		"./internals/ui/html/roles.tmpl",
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

func UserStatPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/users" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./internals/ui/html/admin.panel.tmpl",
		"./internals/ui/html/users.tmpl",
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
