package db

import (
	"context"
	"electro_student/auth/internals/app/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	log "github.com/sirupsen/logrus"
)

type GradesStorage struct {
	databasePool *pgxpool.Pool
}

type userGrade struct {
	TeacherId         int64 `json:"teacher_id" db:"teacher_id"`
	TeacherEmail      string
	TeacherLastName   string
	TeacherFirstName  string
	TeacherMiddleName string
	Discipline        string
	StudentId         int64 `json:"student_id" db:"student_id"`
	StudentEmail      string
	StudentLastName   string
	StudentFirstName  string
	StudentMiddleName string
	Grade             int64
}

func convertJoinedQueryToGrade(input userGrade) models.Grade {
	return models.Grade{
		Teacher: models.User{
			Id:         input.TeacherId,
			Email:      input.TeacherEmail,
			LastName:   input.TeacherLastName,
			FirstName:  input.TeacherFirstName,
			MiddleName: input.TeacherMiddleName,
		},
		Discipline: input.Discipline,
		Student: models.User{
			Id:         input.StudentId,
			Email:      input.StudentEmail,
			LastName:   input.StudentLastName,
			FirstName:  input.StudentFirstName,
			MiddleName: input.StudentMiddleName,
		},
		Grade: input.Grade,
	}
}

func NewGradesStorage(pool *pgxpool.Pool) *GradesStorage {
	storage := new(GradesStorage)
	storage.databasePool = pool
	return storage
}

func (storage *GradesStorage) StudentAndTeacherChecker(grade models.Grade) bool {
	checkerQuery := `SELECT id FROM users WHERE id = $1`

	var teacherId int64

	err := pgxscan.Get(context.Background(), storage.databasePool, &teacherId, checkerQuery, grade.Teacher.Id)

	if err != nil {
		log.Errorln(err)
		return false
	}

	var studentId int64

	err = pgxscan.Get(context.Background(), storage.databasePool, &studentId, checkerQuery, grade.Student.Id)

	if err != nil {
		log.Errorln(err)
		return false
	}

	return true
}

func (storage *GradesStorage) CreateGrade(grade models.Grade) error {
	insertQuery := `INSERT INTO grades(teacher_id, discipline, student_id, grade) VALUES ($1,$2,$3,$4)`

	_, err := storage.databasePool.Exec(context.Background(), insertQuery, grade.Teacher.Id, grade.Discipline, grade.Student.Id, grade.Grade)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}
