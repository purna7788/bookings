package models

import "github.com/purna7788/bookings/cmd/internal/forms"

//Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

//TemplateData holds data sent from handler to renderTemplate
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
	CSRFToken string //Cross-Site Request Forgery Token.
	Error     string
	Flash     string
	Warning   string
	Form      *forms.Form
}
