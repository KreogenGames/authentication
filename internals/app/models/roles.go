package models

type Role struct {
	Id          int64  `json:"id" db:"id"`
	RoleName    string `json:"role_name" db:"role_name"`
	AccessLevel int64  `json:"access_level" db:"access_level"`
}
