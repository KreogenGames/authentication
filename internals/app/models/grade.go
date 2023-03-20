package models

type Grade struct {
	Id         int64  `json:"id" db:"id"`
	Teacher    User   `json:"teacher"`
	Discipline string `json:"discipline" db:"discipline"`
	Student    User   `json:"student"`
	Grade      int64  `json:"grade" db:"grade"`
}
