package handlers

import (
	"electro_student/auth/internals/app/models"
	"electro_student/auth/internals/app/processors"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RolesHandler struct {
	processor *processors.RolesProcessor
}

func NewRolesHandler(processor *processors.RolesProcessor) *RolesHandler {
	handler := new(RolesHandler)
	handler.processor = processor
	return handler
}

func (handler *RolesHandler) AddNewRole(w http.ResponseWriter, r *http.Request) {
	var newRole models.Role

	err := json.NewDecoder(r.Body).Decode(&newRole)

	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.processor.AddNewRole(newRole)

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

func (handler *RolesHandler) FindRoleById(w http.ResponseWriter, r *http.Request) {
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

	user, err := handler.processor.FindRoleById(id)
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

func (handler *RolesHandler) ListRoles(w http.ResponseWriter, r *http.Request) {
	list, err := handler.processor.ListRoles()
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
