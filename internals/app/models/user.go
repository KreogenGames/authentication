package models

type User struct {
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	PhoneNumber string `json:"phoneNumber"`
	Role        int64  `json:"role"`
}
