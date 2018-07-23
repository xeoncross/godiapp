package godiapp

// Our app is so simple there is only a single domain object: a user

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

type UserService interface {
	GetUsers() ([]*User, error)
}
