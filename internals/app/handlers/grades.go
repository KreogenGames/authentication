package handlers

import (
	"electro_student/auth/internals/app/models"
	"electro_student/auth/internals/app/processors"
	"encoding/json"
	"net/http"
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
