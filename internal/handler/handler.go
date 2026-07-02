package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sticky-notes-go-backend/internal/model"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	ctx context.Context
	storer NoteStorer
}

func NewHandler(storer NoteStorer) *Handler {
	return &Handler{
		ctx: context.Background(),
		storer: storer,
	}
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {

	var n NoteRequest

	err := json.NewDecoder(r.Body).Decode(&n)

	if err != nil {
		http.Error(w,"Error Decoind teh incomning request",http.StatusBadRequest)
		return
	}

	note := model.StickyNote{
		Title:   n.Title,
		Content: n.Content,
		 Color:   n.Color,
		Pinned:  n.Pinned,
	}

	res,err := h.storer.CreateNotes(r.Context(),note)
	
	if err != nil {
		http.Error(w,"Error Creating note",http.StatusBadRequest)
		return
	}

	response := NoteResponse{
		ID:        res.ID,
		Title:     res.Title,
		Content:   res.Content,
		 Color:     res.Color,
		 Pinned:    res.Pinned,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func (h *Handler) GetNoteByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r,"id")

	note,err := h.storer.GetNoteById(r.Context(),id)

	if err != nil {
		http.Error(w,"Error Getting note",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&note)
}

func (h *Handler) GetAllNotes(w http.ResponseWriter,r *http.Request) {

	notes,err := h.storer.GettALLNotes(r.Context())

	if err != nil {
		http.Error(w,"Error Getting all notes",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)

}

func (h *Handler) EditNote(w http.ResponseWriter,r *http.Request) {

	var n NoteRequest

	id := chi.URLParam(r,"id")

	err := json.NewDecoder(r.Body).Decode(&n)

	if err != nil {
		http.Error(w,"Error Decoding incoming note request",http.StatusBadRequest)
		return
	}

	note := model.StickyNote{
		Title:   n.Title,
		Content: n.Content,
		Color:   n.Color,
		Pinned:  n.Pinned,
	}

	updated,err := h.storer.EditNote(r.Context(),id,note)

	if err != nil {
		http.Error(w,"Error editing note",http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&updated)

} 


func (h *Handler) DeleteNote(w http.ResponseWriter,r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.storer.DeleteNote(r.Context(),id)

	if err != nil {
		http.Error(w,"Error Deleting note",http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

} 