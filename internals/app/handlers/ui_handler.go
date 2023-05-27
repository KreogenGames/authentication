package handlers

import (
	"electro_student/auth/internals/app/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type templateData struct {
	Roles  []*models.Role
	Users  []*models.User
	Grades []*models.Grade
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

func (handler *GradesHandler) GradesStatPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/grades" {
		http.NotFound(w, r)
		return
	}

	vars := r.URL.Query()

	var studentIdFilter int64 = 0
	if vars.Get("student_id") != "" {
		var err error
		studentIdFilter, err = strconv.ParseInt(vars.Get("student_id"), 10, 64)
		if err != nil {
			WrapError(w, err)
			return
		}
	}
	var teacherIdFilter int64 = 0
	if vars.Get("teacher_id") != "" {
		var err error
		teacherIdFilter, err = strconv.ParseInt(vars.Get("teacher_id"), 10, 64)
		if err != nil {
			WrapError(w, err)
			return
		}
	}
	var gradeFilter int64 = 0
	if vars.Get("grade") != "" {
		var err error
		gradeFilter, err = strconv.ParseInt(vars.Get("grade"), 10, 64)
		if err != nil {
			WrapError(w, err)
			return
		}
	}

	list := handler.processor.SliceGrades(studentIdFilter, teacherIdFilter, strings.Trim(vars.Get("s_email"), "\""), strings.Trim(vars.Get("t_email"), "\""), strings.Trim(vars.Get("discipline"), "\""), gradeFilter)

	data := &templateData{Grades: list}

	files := []string{
		"./internals/ui/html/admin.panel.tmpl",
		"./internals/ui/html/grade.list.tmpl",
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
