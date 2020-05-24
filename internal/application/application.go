package application

import (
	"github.com/go-chi/chi"
	"go.uber.org/fx"

	"mix/internal/application/account"
	"mix/internal/core"
)

var Module = append(
	[]fx.Option{
		fx.Provide(NewApplication),
	},
	account.Module...,
)

type Application struct {
	Account *account.Account
}

func NewApplication(account *account.Account) *Application {
	return &Application{
		Account: account,
	}
}

func (a *Application) Route(router chi.Router) {
	router.Route(core.ApiPrefix, func(r chi.Router) {
		// Client area
		r.Route("/account", a.Account.Route)
	})
}
