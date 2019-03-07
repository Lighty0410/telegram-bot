package database


type Service struct {
	users map[string]User
}

func NewService()*Service{
	service := new(Service)
	service.users = make(map[string]string)
	return service
}
