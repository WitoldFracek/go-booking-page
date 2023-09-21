package handlers

import (
	"net/http"

	"github.com/WitoldFracek/go-bookings-page/pkg/config"
	"github.com/WitoldFracek/go-bookings-page/pkg/models"
	"github.com/WitoldFracek/go-bookings-page/pkg/render"
)

// Repo the tepository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHanflers sets the repository for  the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strings := make(map[string]string)
	strings["test"] = "Hello, again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	strings["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		Strings: strings,
	})
}
