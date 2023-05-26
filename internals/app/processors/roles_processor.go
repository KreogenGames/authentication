package processors

import (
	"electro_student/auth/internals/app/db"
	"electro_student/auth/internals/app/models"
	"errors"
)

type RolesProcessor struct {
	storage *db.RolesStorage
}

func NewRolesProcessor(storage *db.RolesStorage) *RolesProcessor {
	processor := new(RolesProcessor)
	processor.storage = storage
	return processor
}

func (processor *RolesProcessor) AddNewRole(role models.Role) error {
	roleChecker := processor.storage.GetRoleByRoleName(role.RoleName)

	if role.RoleName == "" {
		return errors.New("roleName should not be empty")
	}
	if role.RoleName == roleChecker.RoleName {
		return errors.New("role with this name already exists")
	}
	if role.AccessLevel < 0 {
		return errors.New("accessLevel must be greater than 0")
	}
	if role.AccessLevel >= 10 {
		return errors.New("accessLevel must be less than 10")
	}

	return processor.storage.AddNewRole(role)
}

func (processor *RolesProcessor) FindRoleById(id int64) (models.Role, error) {
	role := processor.storage.GetRoleById(id)

	if role.Id != id {
		return role, errors.New("role not found")
	}

	return role, nil
}

func (processor *RolesProcessor) ListRoles() ([]models.Role, error) {
	return processor.storage.GetRolesList(), nil
}

func (processor *RolesProcessor) SliceRoles() []*models.Role {
	return processor.storage.GetRolesSlice()
}
