package routes

import (
	"sticky-notes-go-backend/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h *handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/notes", func(r chi.Router) {
		
		r.Post("/", h.CreateNote)
		r.Get("/", h.GetAllNotes)
		r.Get("/trash", h.GetTrashNotes)

		r.Get("/{id}", h.GetNoteByID)
		r.Put("/{id}", h.EditNote)
		r.Delete("/{id}", h.DeleteNote)

		r.Post("/{id}/restore", h.RestoreNote)
		r.Delete("/{id}/permanent", h.PermanentDeleteNote)
	})

	return r
}
