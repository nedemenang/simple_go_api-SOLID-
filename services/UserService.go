package services

import (
	// "fmt"
	"errors"

	"github.com/nedemenang/go_authentication_api/interfaces"
	. "github.com/nedemenang/go_authentication_api/models"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) Login(username string, password string) (string, error) {

	// fmt.Println(username)
	user, err := service.GetUserByUserName(username)
	// fmt.Println(user.Password)
	if err != nil {
		return "_", err
	}
	if user.Password != password {
		return "_", errors.New("Authentication failed")
	}

	return user.Username, nil
}

func (service *UserService) RegisterUsers(user UserModel) error {

	err := service.CreateNewUser(user)
	if err != nil {
		return err
	}
	return nil
}
