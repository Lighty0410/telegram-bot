package storage

import "fmt"

// User struct contains basic user's field in Telegram for memdb.
type User struct {
	ID       string
	Token    string
	Password string
}

// ErrUserNotFound returns an error when database cannot find the user.
var ErrUserNotFound = fmt.Errorf("user not found")

type Service interface {
	UpsertUser(u User) error
	GetUser(userID string) (User, error)
}
