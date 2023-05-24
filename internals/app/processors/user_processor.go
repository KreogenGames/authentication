package processors

import (
	"crypto/sha256"
	"electro_student/auth/internals/app/db"
	"electro_student/auth/internals/app/models"
	"errors"
	"fmt"
	"strings"
)

type UsersProcessor struct {
	storage *db.UsersStorage
}

func NewUsersProcessor(storage *db.UsersStorage) *UsersProcessor {
	processor := new(UsersProcessor)
	processor.storage = storage
	return processor
}

const salt = "jhfioejf8653oasdsdakjv4312"

func (processor *UsersProcessor) GeneratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (processor *UsersProcessor) AddNewUser(user models.User) error {
	if user.Email == "" {
		return errors.New("email should not be empty")
	}
	if !strings.ContainsAny(user.Email, "@.") {
		return errors.New("wrong email format")
	}
	if processor.storage.EmailChecker(user.Email) {
		return errors.New("user with this email already exists")
	}
	if user.Role <= 0 {
		return errors.New("missing role id")
	}

	return processor.storage.AddNewUser(user)
}

func (processor *UsersProcessor) CreateUser(user models.User) error {
	if user.Hashed_Pass == "" {
		return errors.New("password should not be empty")
	}
	if user.LastName == "" {
		return errors.New("user's last_name should not be empty")
	}
	if user.FirstName == "" {
		return errors.New("user's first_name should not be empty")
	}
	if !strings.Contains(user.Email, "@edu.mirea.ru") && !strings.Contains(user.Email, "@mirea.ru") {
		return errors.New("email must be in @edu.mirea.ru or @mirea.ru domain")
	}

	user.Hashed_Pass = processor.GeneratePasswordHash(user.Hashed_Pass)

	return processor.storage.CreateUser(user)
}

func (processor *UsersProcessor) FindUser(id int64) (models.User, error) {
	user := processor.storage.GetUserById(id)
	if user.Id != id {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (processor *UsersProcessor) FindUserByEmail(email string) (models.User, error) {
	user := processor.storage.GetUserByEmail(email)

	if user.Email != email {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (processor *UsersProcessor) ListUsers() ([]models.User, error) {
	return processor.storage.GetUsersList(), nil
}

func (processor *UsersProcessor) ListUsersByParams(email string, lastName string, firstName string, middleName string) ([]models.User, error) {
	return processor.storage.GetUsersListByParams(email, lastName, firstName, middleName), nil
}

func (processor *UsersProcessor) UpdateUserPass(email string, newPass string) error {
	newPass = processor.GeneratePasswordHash(newPass)

	err := processor.storage.UpdateUserPass(newPass, email)
	if err != nil {
		return errors.New("password not changed")
	}

	return nil
}
