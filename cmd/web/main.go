package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/taunix/sopcreator/pkg/config"
	"github.com/taunix/sopcreator/pkg/handlers"
	"github.com/taunix/sopcreator/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	// Setup AppConfig
	app.InProduction = false

	// Gather the templates and build Template Cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// Store values in App Config to use in app
	app.TemplateCache = tc
	app.UseCache = false

	// Set the values in handlers and render
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// Show in console that app is starting
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	// Set the server variables
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	// Start the server
	err = srv.ListenAndServe()
	log.Fatal(err)
}
