package models

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
}
