package handler

import (
	"context"
	"time"
	"sticky-notes-go-backend/internal/model"
)

type NoteStorer interface {

	CreateNotes(ctx context.Context, m model.StickyNote) (*model.StickyNote, error)

	GetNoteById(ctx context.Context, id string) (*model.StickyNote, error)
	
	GettALLNotes(ctx context.Context) ([]*model.StickyNote, error)
	
	EditNote(ctx context.Context, id string, data model.StickyNote) (model.StickyNote, error)

	GetTrashNotes(ctx context.Context) ([]*model.StickyNote, error)

	RestoreNote(ctx context.Context, id string) error
	
	PermanentDelete(ctx context.Context, id string) error
	
	DeleteNote(ctx context.Context, id string) error
}

type NoteRequest struct {
	Title   string          `json:"title"`
	Content string          `json:"content"`
	Color   model.NoteColor `json:"color"`
	Pinned  bool            `json:"pinned"`
}

type NoteResponse struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Color     model.NoteColor `json:"color"`
	Pinned    bool            `json:"pinned"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
