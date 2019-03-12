package memdb

// Service struct contains a simple RAM-based database in map-like format.
type Service struct {
	users map[string]User
}

// User struct contains basic user's field in Telegram for memdb.
type User struct {
	ID       string
	Token    string
	Password string
}

// NewService creates database in memory database.
func NewService() *Service {
	return &Service{
		users: make(map[string]User),
	}
}
