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

	// if !processor.storage.StudentAndTeacherChecker(grade.Teacher.Id, grade.Student.Id) {
	// 	return errors.New("teacher or student with such id's not founded")
	// }
	if grade.Discipline == "" {
		return errors.New("discipline should not be empty")
	}
	if grade.Grade <= 0 || grade.Grade > 5 {
		return errors.New("grade must be in range from 0 to 5")
	}

	return processor.storage.CreateGrade(grade)
}

func (processor *GradesProcessor) ListGrades(studentEmailFilter string, disciplineFilter string, gradeFilter int64) ([]models.Grade, error) {
	return processor.storage.GetGradesList(studentEmailFilter, disciplineFilter, gradeFilter), nil
}
