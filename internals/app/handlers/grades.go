package handlers

import (
	"electro_student/auth/internals/app/models"
	"electro_student/auth/internals/app/processors"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type GradesHandler struct {
	processor *processors.GradesProcessor
}

func NewGradesHandler(processor *processors.GradesProcessor) *GradesHandler {
	handler := new(GradesHandler)
	handler.processor = processor
	return handler
}

func (handler *GradesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newGrade models.Grade

	err := json.NewDecoder(r.Body).Decode(&newGrade)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.processor.CreateGrade(newGrade)
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

func (handler *GradesHandler) List(w http.ResponseWriter, r *http.Request) {
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
	var gradeFilter int64 = 0
	if vars.Get("grade") != "" {
		var err error
		gradeFilter, err = strconv.ParseInt(vars.Get("grade"), 10, 64)
		if err != nil {
			WrapError(w, err)
			return
		}
	}

	list, err := handler.processor.ListGrades(studentIdFilter, strings.Trim(vars.Get("discipline"), "\""), gradeFilter)

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
