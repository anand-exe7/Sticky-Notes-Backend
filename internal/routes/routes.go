package routes

import (
	"sticky-notes-go-backend/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(h *handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api/notes", func(r chi.Router) {
		
		r.Post("/", h.CreateNote)
		r.Get("/", h.GetAllNotes)
		r.Get("/trash", h.GetTrashNotes) 

		r.Get("/{id}", h.GetNoteByID)
		r.Put("/{id}", h.EditNote)
		r.Delete("/{id}", h.DeleteNote)

		r.Post("/{id}/restore", h.RestoreNote)
		r.Delete("/{id}/permanent", h.PermanentDeleteNote)
		r.Patch("/{id}/pin", h.TogglePinNote)
		r.Post("/{id}/unlock", h.UnlockNote)
	})

	return r
}
