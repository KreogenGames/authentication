package models

type User struct {
	Id          int64  `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Hashed_Pass string `json:"hashed_pass" db:"hashed_pass"`
	LastName    string `json:"lastName" db:"lastName"`
	FirstName   string `json:"firstName" db:"firstName"`
	MiddleName  string `json:"middleName" db:"middleName"`
	PhoneNumber string `json:"phoneNumber" db:"phoneNumber"`
	Role        int64  `json:"role" db:"role"`
	//Role        Role   `json:"role"`
}
