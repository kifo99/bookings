package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kifo99/bookings/pkg/config"
	"github.com/kifo99/bookings/pkg/models"
	"github.com/kifo99/bookings/pkg/render"
)

//Repo the repository used by handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//reciver//
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Reservation renders make reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

//Renders the room page
func (m *Repository) Penthouse(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "penthouse.page.tmpl", &models.TemplateData{})
}

//Renders the room page
func (m *Repository) Cheep(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "cheep.page.tmpl", &models.TemplateData{})
}

//Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and enda date is %s", start, end)))
}

type jsonRespons struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and send JSON respons
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonRespons{
		OK:      false,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

//Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
