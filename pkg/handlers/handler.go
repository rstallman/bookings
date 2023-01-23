package handlers

import (
	"net/http"

	"github.com/rstallman/bookings/pkg/config"
	"github.com/rstallman/bookings/pkg/models"
	"github.com/rstallman/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	StringMap := map[string]string{
		"test": "Hello again",
	}

	remoteIP := m.App.Session.Get(r.Context(), "remote_ip")
	if remoteIP == nil {
		remoteIP = ""
	}
	StringMap["remote_ip"] = remoteIP.(string)
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: StringMap})
}
