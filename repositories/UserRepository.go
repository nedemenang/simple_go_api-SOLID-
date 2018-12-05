package repositories

import (
	// "fmt"
	"log"

	"github.com/afex/hystrix-go/hystrix" // I dont know
	"github.com/nedemenang/go_authentication_api/interfaces"
	"github.com/nedemenang/go_authentication_api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// type UserDAO struct {
// 	Server   string
// 	Database string
// }

type UserRepositoryWithCircuitBreaker struct {
	UserRepository interfaces.IUserRepository
}

func (repository *UserRepositoryWithCircuitBreaker) GetUserByUserName(name string) (models.UserModel, error) {

	output := make(chan models.UserModel, 1)
	hystrix.ConfigureCommand("get_user_by_username", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_user_by_username", func() error {
		user, _ := repository.UserRepository.GetUserByUserName(name)
		// fmt.Println(user)
		output <- user
		
		return nil

	}, nil)
	select {
		case out := <-output:
			// fmt.Println(out)
			return out, nil
		case err := <-errors:
			println(err)
			return models.UserModel{}, err
		// default:
		// 	// fmt.Println("I am here")
		// 	// fmt.Println(u)
		// return u, nil
	}
}

func (repository *UserRepositoryWithCircuitBreaker) GetUserByEmail(email string) (models.UserModel, error) {

	output := make(chan models.UserModel, 1)

	hystrix.ConfigureCommand("get_user_by_email", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_user_by_email", func() error {
		user, _ := repository.UserRepository.GetUserByUserName(email)

		output <- user
		return nil

	}, nil)
	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.UserModel{}, err
	default:
		return models.UserModel{}, nil
	}
}

func (repository *UserRepositoryWithCircuitBreaker) CreateNewUser(user models.UserModel) error {
	// fmt.Println("Im here")

	hystrix.ConfigureCommand("create_new_user", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("create_new_user", func() error {
		repository.UserRepository.CreateNewUser(user)

		return nil

	}, nil)
	select {
	case err := <-errors:
		println(err)
		return err
	default:
		return nil
	}
}

func (repository *UserRepositoryWithCircuitBreaker) ResetPassword(user models.UserModel) error {

	hystrix.ConfigureCommand("reset_password", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("reset_password", func() error {
		repository.UserRepository.ResetPassword(user)

		return nil

	}, nil)
	select {
	case err := <-errors:
		println(err)
		return err
	default:
		return nil
	}
}

func (repository *UserRepositoryWithCircuitBreaker) DeleteUser(user models.UserModel) error {

	hystrix.ConfigureCommand("delete_user", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("delete_user", func() error {
		repository.UserRepository.DeleteUser(user)

		return nil

	}, nil)
	select {
	case err := <-errors:
		println(err)
		return err
	default:
		return nil
	}
}

type UserRepository struct {
	// interfaces.IDbHandler
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

// var config = Config{}
// var dao = UserDAO{}

// Establish a connection to database
func (m *UserRepository) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (repository *UserRepository) GetUserByUserName(name string) (models.UserModel, error) {
	var user models.UserModel
	err := db.C(COLLECTION).Find(bson.M{"username": name}).One(&user)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(user.Email)
	return user, err

}

func (repository *UserRepository) GetUserByEmail(email string) (models.UserModel, error) {
	var user models.UserModel
	err := db.C(COLLECTION).Find(bson.M{"email": email}).One(&user)
	return user, err
}

func (repository *UserRepository) CreateNewUser(user models.UserModel) error {

	err := db.C(COLLECTION).Insert(&user)
	if err != nil {
		log.Fatal(err)
	}
	return nil

}

func (repository *UserRepository) ResetPassword(user models.UserModel) error {

	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (repository *UserRepository) DeleteUser(user models.UserModel) error {

	err := db.C(COLLECTION).Remove(&user)
	if err != nil {
		log.Fatal(err)
	}
	return nil

}
