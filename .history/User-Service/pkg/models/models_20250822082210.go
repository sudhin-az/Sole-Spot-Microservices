package models

type UserLogin struct {
	Email    string
	Password string
}

type UserDetails struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
}

type UserSignUp struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
}

type UserDetail struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
}
