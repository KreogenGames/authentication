package processors

import (
	"electro_student/auth/internals/app/db"
	"electro_student/auth/internals/app/models"
	"errors"
)

type GradesProcessor struct {
	storage *db.GradesStorage
}

func NewGradesProcessor(storage *db.GradesStorage) *GradesProcessor {
	processor := new(GradesProcessor)
	processor.storage = storage
	return processor
}

func (processor *GradesProcessor) CreateGrade(grade models.Grade) error {
	if grade.Teacher.Id <= 0 {
		return errors.New("teacher id shall be filled")
	}

	if grade.Discipline == "" {
		return errors.New("discipline should not be empty")
	}

	if grade.Student.Id <= 0 {
		return errors.New("student id shall be filled")
	}

	if grade.Grade <= 0 || grade.Grade >= 5 {
		return errors.New("grade must be in range from 0 to 5")
	}

	return processor.storage.CreateGrade(grade)
}
