package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sticky-notes-go-backend/internal/model"
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
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)

}