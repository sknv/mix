package controllers

import (
	"net/http"

	"github.com/go-chi/chi"

	"mix/pkg/web/response"
)

type AccountController struct{}

func NewAccountController() *AccountController {
	return &AccountController{}
}

func (a *AccountController) Route(router chi.Router) {
	router.Get("/hello", a.Hello)
}

func (a *AccountController) Hello(w http.ResponseWriter, r *http.Request) {
	response.RenderData(w, r, map[string]string{"hello": "world"})
}
