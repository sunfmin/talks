package main

type User struct {
	Email string
	Name  string
}

type Service interface {
	Login(email string, password string) (user *User, err error)
	Register(email string, password string, passwordConfirmation string) (user *User, err error)
}

type AuthService interface {
	MyUnreadCount() (count int)
}

// endmain OMIT
