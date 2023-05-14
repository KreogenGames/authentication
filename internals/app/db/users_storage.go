package db

import (
	"context"
	"electro_student/auth/internals/app/models"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type UsersStorage struct {
	databasePool *pgxpool.Pool
}

func NewUsersStorage(pool *pgxpool.Pool) *UsersStorage {
	storage := new(UsersStorage)
	storage.databasePool = pool
	return storage
}

func (storage *UsersStorage) AddNewUser(newUser models.User) error {
	ctx := context.Background()
	tx, err := storage.databasePool.Begin(ctx)
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
	}()

	searchQuery := `SELECT email FROM users WHERE email = $1`

	email := ""

	err = pgxscan.Get(ctx, tx, &email, searchQuery, newUser.Email)

	if err != nil {
		log.Errorln(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
		return err
	}

	insertQuery := `INSERT INTO users(email, role) VALUES ($1, $2)`

	_, err = tx.Exec(ctx, insertQuery, newUser.Email, newUser.Role)

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

func (storage *UsersStorage) CreateUser(user models.User) error {
	// query := "INSERT INTO users (password, lastName, firstName, middleName, phoneNumber) VALUES ($1, $2, $3, $4, $5) WHERE email = $6"
	query := `UPDATE users SET "hashed_pass" = $1, "lastName" = $2, "firstName" = $3, "middleName" = $4, "phoneNumber" = $5 WHERE "email" = $6`

	_, err := storage.databasePool.Exec(context.Background(), query, user.Hashed_Pass, user.LastName, user.FirstName, user.MiddleName, user.PhoneNumber, user.Email)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}

func (storage *UsersStorage) GetUserById(id int64) models.User {
	query := `SELECT id, "email", "hashed_pass", "lastName", "firstName", "middleName", "phoneNumber" FROM users WHERE id = $1`

	var result models.User

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, id)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (storage *UsersStorage) GetUserByEmail(email string) models.User {
	query := `SELECT id, "email", "hashed_pass", "lastName", "firstName", "middleName", "phoneNumber" FROM users WHERE "email" = $1`

	var result models.User

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, email)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (storage *UsersStorage) GetUsersList(email string, lastName string, firstName string, middleName string) []models.User {
	query := `SELECT id, "email", "hashed_pass", "lastName", "firstName", "middleName", "phoneNumber" FROM users WHERE 1=1`

	placeholderNum := 1
	args := make([]interface{}, 0)

	if email != "" {
		query += fmt.Sprintf(` AND "email" ILIKE $%d`, placeholderNum)
		args = append(args, fmt.Sprintf("%%%s%%", email))
		placeholderNum++
	}
	if lastName != "" {
		query += fmt.Sprintf(` AND "lastName" ILIKE $%d`, placeholderNum)
		args = append(args, fmt.Sprintf("%%%s%%", lastName))
		placeholderNum++
	}
	if firstName != "" {
		query += fmt.Sprintf(` AND "firstName" ILIKE $%d`, placeholderNum)
		args = append(args, fmt.Sprintf("%%%s%%", firstName))
		placeholderNum++
	}
	if middleName != "" {
		query += fmt.Sprintf(` AND "middleName" ILIKE $%d`, placeholderNum)
		args = append(args, fmt.Sprintf("%%%s%%", middleName))
		placeholderNum++
	}

	var result []models.User

	err := pgxscan.Select(context.Background(), storage.databasePool, &result, query, args...)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (storage *UsersStorage) UpdateUserPass(newPass string, email string) error {
	ctx := context.Background()
	tx, err := storage.databasePool.Begin(ctx)

	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
	}()

	updateQuery := `UPDATE users SET "hashed_pass" = $1 WHERE "email" = $2`

	_, err = tx.Exec(ctx, updateQuery, newPass, email)

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
