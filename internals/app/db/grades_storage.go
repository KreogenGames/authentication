package db

import (
	"context"
	"electro_student/auth/internals/app/models"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	log "github.com/sirupsen/logrus"
)

type GradesStorage struct {
	databasePool *pgxpool.Pool
}

type userGrade struct {
	TeacherId         int64 `json:"teacher_id" db:"userid"`
	TeacherEmail      string
	TeacherLastName   string
	TeacherFirstName  string
	TeacherMiddleName string
	Discipline        string
	StudentId         int64 `json:"student_id" db:"gradeid"`
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

func (storage *UsersStorage) StudentAndTeacherChecker(teacher_id int64, student_id int64) bool {
	checkerQuery := `SELECT id FROM users WHERE id = $1`

	var teacherId int64

	err := pgxscan.Get(context.Background(), storage.databasePool, &teacherId, checkerQuery, teacher_id)

	if err != nil {
		log.Errorln(err)
		return false
	}

	var studentId int64

	err = pgxscan.Get(context.Background(), storage.databasePool, &studentId, checkerQuery, student_id)

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

func (storage *GradesStorage) GetGradesList(studentEmailFilter string, disciplineFilter string, gradeFilter int64) []models.Grade {
	query := `SELECT users.id AS userid, users.email, users.lastName, users.firstName, users.middleName, grades.id AS gradeid, grades.discipline, grades.grade FROM users JOIN grades g on users.id = g.student_id WHERE 1=1`

	placeholderNum := 1
	args := make([]interface{}, 0)

	if studentEmailFilter != "" {
		query += fmt.Sprintf(" AND users.id = $%d", placeholderNum)
		args = append(args, studentEmailFilter)
		placeholderNum++
	}
	if disciplineFilter != "" {
		query += fmt.Sprintf(" AND discipline ILIKE $%d", placeholderNum)
		args = append(args, disciplineFilter)
		placeholderNum++
	}
	if gradeFilter != 0 {
		query += fmt.Sprintf(" AND grade ILIKE $%d", placeholderNum)
		args = append(args, gradeFilter)
		placeholderNum++
	}

	var dbResult []userGrade

	err := pgxscan.Get(context.Background(), storage.databasePool, &dbResult, query, args...)
	if err != nil {
		log.Errorln(err)
	}

	result := make([]models.Grade, len(dbResult))

	for idx, dbEntity := range dbResult {
		result[idx] = convertJoinedQueryToGrade(dbEntity)
	}

	return result
}
