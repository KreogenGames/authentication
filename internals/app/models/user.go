package models

type User struct {
	Id          int64  `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Hashed_Pass string `json:"hashed_pass" db:"hashed_pass"`
	LastName    string `json:"lastName" db:"last_name"`
	FirstName   string `json:"firstName" db:"first_name"`
	MiddleName  string `json:"middleName" db:"middle_name"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	Role        int64  `json:"role" db:"role"`
	//Role        Role   `json:"role"`
}
