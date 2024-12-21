package scribe

import (
	"context"
	"log"
	"tpot_scribe/pkg/convert"
	"tpot_scribe/pkg/ui/templates"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var version = "0.0.0"

// App struct
type App struct {
	ctx    context.Context
	Router chi.Router
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Router: NewRouter(),
	}
}

// startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) DomReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) Greet() string {
	return "hello there!"
}

func (a *App) LoadDocument(filepath string) string {
	html, err := convert.DocxToHTML(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return html
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/initial", templ.Handler(templates.Pages([]struct {
		Path  string
		Label string
	}{
		{"/greet", "Greet form"},
	}, struct {
		Version string
		Text    string
	}{
		version, "No update available",
	})).ServeHTTP)
	r.Get("/greet", templ.Handler(templates.GreetForm("/greet")).ServeHTTP)
	r.Post("/greet", templates.Greet)
	r.Get("/modal", templ.Handler(templates.TestPage("#modal", "outerHTML")).ServeHTTP)
	r.Post("/modal", templ.Handler(templates.ModalPreview("Title for the modal", "Sample Data")).ServeHTTP)
	return r
}
