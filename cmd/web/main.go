package main

import (
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
	"ullibniss/pkg/config"
	"ullibniss/pkg/handlers"
	"ullibniss/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

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
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	//_ = http.ListenAndServe(":8080", nil)

	portNumber := ":8080"

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

