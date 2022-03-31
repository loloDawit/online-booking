package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"gitlab.nordstrom.com/online-booking/pkg/config"
	"gitlab.nordstrom.com/online-booking/pkg/handlers"
	"gitlab.nordstrom.com/online-booking/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	// update for production
	app.IsProd = false

	// Initialize a new session manager and configure the session lifetime.
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.IsProd

	tc, err := render.CreatTemplateCache()
	if err != nil {
		log.Fatal("Can not create a template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf(fmt.Sprintf("Server running on port %s", portNumber))
	err = server.ListenAndServe()

	log.Fatal(err)
}
