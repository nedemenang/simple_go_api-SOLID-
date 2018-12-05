package main

import (
	"net/http"

	"github.com/nedemenang/go_authentication_api/router"
)

func main() {

	http.ListenAndServe(":8080", router.ChiRouter().InitRouter())
}
