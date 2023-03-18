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

func (storage *UsersStorage) CreateUser(user models.User) error {
	// query := "INSERT INTO users (password, lastName, firstName, middleName, phoneNumber) VALUES ($1, $2, $3, $4, $5) WHERE email = $6"
	query := "UPDATE users SET password = $1, lastName = $2, firstName = $3, middleName = $4, phoneNumber = $5 WHERE email = $6"

	_, err := storage.databasePool.Exec(context.Background(), query, user.Password, user.LastName, user.FirstName, user.MiddleName, user.PhoneNumber, user.Email)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}

func (storage *UsersStorage) GetUserById(id int64) models.User {
	query := "SELECT * FROM user WHERE id = $1"

	var result models.User

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, id)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (storage *UsersStorage) GetUserByEmail(email string) models.User {
	query := "SELECT * FROM user WHERE email = $1"

	var result models.User

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, email)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (storage *UsersStorage) GetUsersListByLastName(lastNameFilter string) []models.User {
	query := "SELECT email, lastName, firstName, middleName, phoneNumber FROM users"

	args := make([]interface{}, 0)

	if lastNameFilter != "" {
		query += " WHERE lastName LIKE $1"
		args = append(args, fmt.Sprintf("%%%s%%", lastNameFilter))
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

	updateQuery := "UPDATE users SET password = $1 WHERE email = $2"

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
