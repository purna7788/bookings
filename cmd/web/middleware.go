package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WritToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" I am middleware func")
		next.ServeHTTP(w, r)

	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(
		http.Cookie{
			HttpOnly: true,
			Path:     "/",
			Secure:   app.IsProduction,
			SameSite: http.SameSiteLaxMode,
		})

	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)

}
