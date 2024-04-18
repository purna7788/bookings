package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/purna7788/bookings/cmd/internal/config"
	"github.com/purna7788/bookings/cmd/internal/handlers"
	"github.com/purna7788/bookings/cmd/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//to initialize template cache

	var err error

	app.IsProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	tc, err := render.CreateCache()
	if err != nil {
		log.Fatal(" fatal error ")
	}
	app.TemplateCache = tc
	app.UseCache = false
	app.Session = session

	r := handlers.NewRepository(&app)

	handlers.NewHandler(&r)

	//initializing global tc variable in render package
	render.NewTemplate(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	//http.ListenAndServe(":8080",nil)

	srv := http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}
