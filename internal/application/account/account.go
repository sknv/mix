package account

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/fx"

	"mix/pkg/web/response"
)

var Module = []fx.Option{
	fx.Provide(NewAccount),
}

type Account struct {
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) Route(router chi.Router) {
	router.Get("/hello", a.Hello)
}

func (a *Account) Hello(w http.ResponseWriter, r *http.Request) {
	response.RenderData(w, r, map[string]string{"hello": "world"})
}
