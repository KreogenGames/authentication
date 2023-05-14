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

func (proccessor *UsersProcessor) AddNewUser(user models.User) error {
	if user.Email == "" {
		return errors.New("email should not be empty")
	}
	if strings.Contains(user.Email, "@[a-z].[a-z]") == false {
		return errors.New("wrong email format")
	}
	if user.Role < 0 || user.Role > 2 {
		return errors.New("missing role id")
	}

	return proccessor.storage.AddNewUser(user)
}

func (processor *UsersProcessor) CreateUser(user models.User) error {
	if user.Hashed_Pass == "" {
		return errors.New("password should not be empty")
	}
	if user.LastName == "" {
		return errors.New("user's lastName should not be empty")
	}
	if user.FirstName == "" {
		return errors.New("user's firstName should not be empty")
	}
	if strings.Contains(user.Email, "@edu.mirea.ru") == false && strings.Contains(user.Email, "@mirea.ru") == false {
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

func (processor *UsersProcessor) ListUsers(email string, lastName string, firstName string, middleName string) ([]models.User, error) {
	return processor.storage.GetUsersList(email, lastName, firstName, middleName), nil
}

func (processor *UsersProcessor) UpdateUserPass(email string, newPass string) error {
	newPass = processor.GeneratePasswordHash(newPass)

	err := processor.storage.UpdateUserPass(newPass, email)
	if err != nil {
		return errors.New("password not changed")
	}

	return nil
}
