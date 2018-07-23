package http

import (
	"fmt"
	"net/http"

	"bitbucket.org/xeoncross/godiapp"
	"github.com/julienschmidt/httprouter"
)

// Handles actuall HTTP endpoints

func NewRouter(us godiapp.UserService) *httprouter.Router {
	h := Handlers{us}

	router := httprouter.New()
	router.GET("/", h.indexHandler)
	router.ServeFiles("/static/*filepath", http.Dir("./static"))
	return router
}

type Handlers struct {
	db godiapp.UserService
}

// Only this one route is defined
func (h *Handlers) indexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := h.db.GetUsers()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "URL: %s\n", r.URL)

	for _, u := range users {
		fmt.Fprintf(w, "User: %+v\n", u)
	}
}

// TODO Input Validation: https://husobee.github.io/golang/validation/2016/01/08/input-validation.html
