package postgres

import (
	"errors"
	"github.com/abhicit0802/design-golang-errors/app"
)

type DbUserService struct {
}

func (us DbUserService) FindById(id int) (*app.User, error) {
	if id == 1 {
		return nil, &app.Error{
			Code:      app.ENOTFOUND,
			Message:   "User not found for id 1",
			Operation: "DbUserService.FindById",
			Err:       errors.New("user not found"),
		}
	}
	return &app.User{Id: id, Name: "Amit", Email: "foo@bar"}, nil
}

func (us DbUserService) Create(user app.User) error {
	err := us.insert(user)
	if err != nil {
		return &app.Error{
			Code:      app.ECONFLICT,
			Message:   "Conflict in insert",
			Operation: "UserService.createUser",
			Err:       err,
		}
	}
	err = us.attachRole(user)
	if err != nil {
		return &app.Error{
			Code:      app.ECONFLICT,
			Message:   "Conflict in attach role",
			Operation: "UserService.createUser",
			Err:       err,
		}
	}
	return nil
}

func (us DbUserService) insert(user app.User) error {
	if user.Id == 1 {
		return &app.Error{
			Code:      app.ECONFLICT,
			Message:   "Conflicted user insert",
			Operation: "app.DbUserService.insert",
			Err:       errors.New("insert error on id 1"),
		}
	}
	return nil
}

func (us DbUserService) attachRole(user app.User) error {
	if user.Id == 1 {
		return &app.Error{
			Code:      app.ECONFLICT,
			Message:   "Conflicted user attachRole",
			Operation: "app.DbUserService.attachRole",
			Err:       errors.New("attach role error for id 1"),
		}
	}
	return nil
}
