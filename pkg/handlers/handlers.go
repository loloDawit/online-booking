package handlers

import (
	"net/http"

	"gitlab.nordstrom.com/online-booking/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.html")
}
