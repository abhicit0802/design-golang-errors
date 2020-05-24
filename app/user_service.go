package app

type UserService interface {
	FindById(id int) (*User, error)
	Create(user User) error
}
