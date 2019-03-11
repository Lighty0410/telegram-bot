package memdb

type Service struct {
	users map[string]User
}

type User struct {
	ID       string
	Token    string
	Password string
}

func NewService() *Service {
	return &Service{
		users: make(map[string]User),
	}
}
