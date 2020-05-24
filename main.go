package main

import (
	"fmt"
	"github.com/abhicit0802/design-golang-errors/app"
	"github.com/abhicit0802/design-golang-errors/postgres"
)

func main() {
	var userService app.UserService
	userService = postgres.DbUserService{}

	_, err := userService.FindById(1)
	if err != nil && app.ErrorCode(err) == app.ENOTFOUND {
		fmt.Println(err.Error())
	}

	user, err := userService.FindById(2)
	if err != nil && app.ErrorCode(err) == app.ENOTFOUND {
		fmt.Println("Not found")
	}
	user.Id = 1

	err = userService.Create(*user)
	if msg := app.ErrorMessage(err); err != nil && msg != "" {
		fmt.Println(msg)
	}
	if code := app.ErrorCode(err); err != nil && code != "" {
		fmt.Println(code)
		fmt.Println(err.Error())
	}

}
