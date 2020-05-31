package api

import (
	"github.com/go-chi/chi"

	"mix/app/api/controllers"
)

type Application struct {
	AccountController *controllers.AccountController
}

func NewApplication(
	accountController *controllers.AccountController,
) *Application {
	return &Application{
		AccountController: accountController,
	}
}

func (a *Application) Route(router chi.Router) {
	router.Route("/account", a.AccountController.Route)
}
