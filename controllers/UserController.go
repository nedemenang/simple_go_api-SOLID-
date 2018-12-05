package controllers

import (
	// "strings"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nedemenang/go_authentication_api/interfaces"
	"github.com/nedemenang/go_authentication_api/models"
	"github.com/nedemenang/go_authentication_api/repositories"
	"github.com/thedevsaddam/renderer" // "github.com/go-chi/chi"
	"gopkg.in/mgo.v2/bson"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type UserController struct {
	interfaces.IUserService
}

var rnd *renderer.Render
var dao = repositories.UserRepository{}

func (controller *UserController) Register(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var u models.UserModel
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if u.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if u.Firstname == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if u.Lastname == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if u.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if u.Username == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	u.ID = bson.NewObjectId()

	err := controller.RegisterUsers(u)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	respondWithJson(w, http.StatusCreated, u)

}

func (controller *UserController) Authenticate(res http.ResponseWriter, req *http.Request) {

	var u models.UserModel
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		respondWithError(res, http.StatusProcessing, "Invalid error")
		return

		// rnd.JSON(res, http.StatusProcessing, err)
		// return
	}

	if u.Password == "" {
		respondWithError(res, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if u.Username == "" {
		respondWithError(res, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := controller.Login(u.Username, u.Password)

	fmt.Println(result)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid request payload")
		return
	}

	respondWithJson(res, http.StatusOK, result+" is successfully logged in")

}

func init() {
	dao.Server = "localhost"
	dao.Database = "go_users"
	dao.Connect()
}
