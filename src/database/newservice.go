package database

type Service struct {
	users map[string]User
}

type User struct {
	Token    string
	Password string
}

func NewService() *Service {
	service := new(Service)
	service.users = make(map[string]User)
	return service
}
