package servicecontainer


import (
	"sync"

	"github.com/nedemenang/go_authentication_api/controllers"
	"github.com/nedemenang/go_authentication_api/repositories"
	"github.com/nedemenang/go_authentication_api/services"
)

type IServiceContainer interface {
	InjectUserController() controllers.UserController
}

type kernel struct{}

func (k *kernel) InjectUserController() controllers.UserController {

	// sqlConn, _ := sql.Open("sqlite3", "/var/tmp/user.db")
	// sqliteHandler := &infrastructures.SQLiteHandler{}
	// sqliteHandler.Conn = sqlConn

	userRepository := &repositories.UserRepository{}
	userService := &services.UserService{&repositories.UserRepositoryWithCircuitBreaker{userRepository}}
	userController := controllers.UserController{userService}

	return userController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
