package db

import (
	"context"
	"electro_student/auth/internals/app/models"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	log "github.com/sirupsen/logrus"
)

type GradesStorage struct {
	databasePool *pgxpool.Pool
}

// type userGrade struct {
// 	TeacherId         int64 `json:"teacher_id"`
// 	TeacherEmail      string
// 	TeacherLastName   string
// 	TeacherFirstName  string
// 	TeacherMiddleName string
// 	Discipline        string
// 	StudentId         int64 `json:"student_id"`
// 	StudentEmail      string
// 	StudentLastName   string
// 	StudentFirstName  string
// 	StudentMiddleName string
// 	Grade             int64
// }

// func convertJoinedQueryToGrade(input userGrade) models.Grade {
// 	return models.Grade{
// 		Teacher: models.User{
// 			Id:         input.TeacherId,
// 			Email:      input.TeacherEmail,
// 			LastName:   input.TeacherLastName,
// 			FirstName:  input.TeacherFirstName,
// 			MiddleName: input.TeacherMiddleName,
// 		},
// 		Discipline: input.Discipline,
// 		Student: models.User{
// 			Id:         input.StudentId,
// 			Email:      input.StudentEmail,
// 			LastName:   input.StudentLastName,
// 			FirstName:  input.StudentFirstName,
// 			MiddleName: input.StudentMiddleName,
// 		},
// 		Grade: input.Grade,
// 	}
// }

func NewGradesStorage(pool *pgxpool.Pool) *GradesStorage {
	storage := new(GradesStorage)
	storage.databasePool = pool
	return storage
}

func (storage *GradesStorage) CreateGrade(grade models.Grade) error {
	ctx := context.Background()
	tx, err := storage.databasePool.Begin(ctx)

	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
	}()

	query := "SELECT id FROM users WHERE id = $1"

	id := -1

	err = pgxscan.Get(ctx, tx, &id, query, grade.Teacher.Id)

	if err != nil {
		log.Errorln(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
		return err
	}

	if id == -1 {
		return errors.New("user not found")
	}

	insertQuery := `INSERT INTO grades(teacher_id, discipline, student_id, grade) VALUES ($1,$2,$3,$4)`

	_, err = tx.Exec(ctx, insertQuery, grade.Teacher.Id, grade.Discipline, grade.Student.Id, grade.Grade)

	if err != nil {
		log.Errorln(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
		return err
	}
	err = tx.Commit(context.Background())

	if err != nil {
		log.Errorln(err)
	}

	return err
}
