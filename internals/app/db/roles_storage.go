package db

import (
	"context"
	"electro_student/auth/internals/app/models"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type RolesStorage struct {
	databasePool *pgxpool.Pool
}

func NewRolesStorage(pool *pgxpool.Pool) *RolesStorage {
	storage := new(RolesStorage)
	storage.databasePool = pool
	return storage
}

func (storage *RolesStorage) AddNewRole(role models.Role) error {
	insertQuery := `INSERT INTO roles (role_name, access_level) VALUES ($1, $2)`

	role.RoleName = strings.ToLower(role.RoleName)

	_, err := storage.databasePool.Exec(context.Background(), insertQuery, role.RoleName, role.AccessLevel)

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}

func (storage *RolesStorage) GetRoleById(id int64) models.Role {
	query := `SELECT * FROM roles WHERE id = $1`

	var result models.Role

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, id)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (storage *RolesStorage) GetRoleByRoleName(roleName string) models.Role {
	query := `SELECT * FROM roles WHERE role_name = $1`

	var result models.Role

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, roleName)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (storage *RolesStorage) GetRolesList() []models.Role {
	query := `SELECT * FROM roles`

	var result []models.Role

	err := pgxscan.Select(context.Background(), storage.databasePool, &result, query)

	if err != nil {
		log.Errorln(err)
	}

	return result
}

func (storage *RolesStorage) GetRolesSlice() []*models.Role {
	query := `SELECT * FROM roles`

	var result []models.Role

	err := pgxscan.Select(context.Background(), storage.databasePool, &result, query)

	if err != nil {
		log.Errorln(err)
	}

	var roleSlice []*models.Role
	for i := range result {
		roleSlice = append(roleSlice, &result[i])
	}

	return roleSlice
}
