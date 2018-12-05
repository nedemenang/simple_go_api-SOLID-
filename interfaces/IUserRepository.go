package interfaces

import (
	"github.com/nedemenang/go_authentication_api/models"
)

// repository for database actions
type IUserRepository interface {

	GetUserByUserName(name string) (models.UserModel, error)
	
	GetUserByEmail(email string) (models.UserModel, error)

	CreateNewUser(user models.UserModel)  error

	ResetPassword(user models.UserModel) error

	DeleteUser(user models.UserModel) error

}