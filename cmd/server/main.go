package main

import (
	"fmt"
	"log"
	"net/http"

	"gitlab.nordstrom.com/online-booking/pkg/config"
	"gitlab.nordstrom.com/online-booking/pkg/handlers"
	"gitlab.nordstrom.com/online-booking/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

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
