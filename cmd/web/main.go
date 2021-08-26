package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/UdemyStudies21/Bookings/pkg/config"
	"github.com/UdemyStudies21/Bookings/pkg/handlers"
	"github.com/UdemyStudies21/Bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main is the main function
func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	svr := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = svr.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
