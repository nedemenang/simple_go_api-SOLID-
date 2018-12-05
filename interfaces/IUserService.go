package interfaces

import (
	"github.com/nedemenang/go_authentication_api/models"
)

type IUserService interface {
	Login(username string, password string) (string, error)

	RegisterUsers(user models.UserModel) error
}
