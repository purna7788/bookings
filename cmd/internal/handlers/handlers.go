package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/purna7788/bookings/cmd/internal/config"
	"github.com/purna7788/bookings/cmd/internal/models"
	"github.com/purna7788/bookings/cmd/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	stringMap["ip_address"] = m.app.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, r, "about.page.tmpl", models.TemplateData{StringMap: stringMap})
}

// Generals renders room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", models.TemplateData{})
}

// Majors renders room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", models.TemplateData{})
}

// Reservation renders the make reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))

}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJson handles request for availability and sends JSON data
func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	res := jsonResponse{
		Ok:      true,
		Message: "Available",
	}
	resJson, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", models.TemplateData{})
}
