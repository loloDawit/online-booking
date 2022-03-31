package handlers

import (
	"net/http"

	"gitlab.nordstrom.com/online-booking/pkg/config"
	"gitlab.nordstrom.com/online-booking/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.html")
}

func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.html")
}
