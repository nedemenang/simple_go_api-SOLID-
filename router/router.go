package router

import (
	"sync"

	"github.com/go-chi/chi"
	"github.com/nedemenang/go_authentication_api/servicecontainer"
)

// IChiRouter is a something
type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	userController := servicecontainer.ServiceContainer().InjectUserController()

	r := chi.NewRouter()

	r.HandleFunc("/authenticate", userController.Authenticate)
	r.HandleFunc("/register", userController.Register)
	return r

}

var (
	m          *router
	routerOnce sync.Once
)

// ChiRouter i dont know
func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
