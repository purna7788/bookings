package handlers

import (
	"net/http"

	"github.com/purna7788/bookings/pkg/config"
	"github.com/purna7788/bookings/pkg/models"
	"github.com/purna7788/bookings/pkg/render"
)

type Repository struct {
	app *config.AppConfig
}

var Repo *Repository

func NewRepository(a *config.AppConfig) Repository {
	return Repository{app: a}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIPaddress := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote_ip", remoteIPaddress)
	render.RenderTemplate(w, "home.page.tmpl", models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	stringMap["ip_address"] = m.app.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.tmpl", models.TemplateData{StringMap: stringMap})
}
