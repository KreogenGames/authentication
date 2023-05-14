package handlers

import (
	"electro_student/auth/internals/app/models"
	"electro_student/auth/internals/app/processors"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type UsersHandler struct {
	processor *processors.UsersProcessor
}

func NewUsersHandler(processor *processors.UsersProcessor) *UsersHandler {
	handler := new(UsersHandler)
	handler.processor = processor
	return handler
}

func (handler *UsersHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.processor.AddNewUser(newUser)

	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   "",
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.processor.CreateUser(newUser)

	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   "",
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	user, err := handler.processor.FindUser(id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   user,
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) FindUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["email"])
	if vars["email"] == "" {
		WrapError(w, errors.New("missing email"))
		return
	}

	email := string(vars["email"])

	user, err := handler.processor.FindUserByEmail(email)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   user,
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) List(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	list, err := handler.processor.ListUsers(strings.Trim(vars.Get("email"), "\""), strings.Trim(vars.Get("lastName"), "\""), strings.Trim(vars.Get("firstName"), "\""), strings.Trim(vars.Get("middleName"), "\""))

	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)
}
