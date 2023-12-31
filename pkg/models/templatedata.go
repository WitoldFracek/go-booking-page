package models

// TemplateData holds dat sent from handlers to templates
type TemplateData struct {
	Strings   map[string]string
	Ints      map[string]int
	Floats    map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}