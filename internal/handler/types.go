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

	TogglePin(ctx context.Context, id string) error
}

type NoteRequest struct {
	Title          string                `json:"title"`
	Content         string                `json:"content"`
	Color          model.NoteColor       `json:"color"`
	Pinned         bool                  `json:"pinned"`
	Category        string                `json:"category"`
	IsLocked        bool                  `json:"isLocked"`
	Password       string                `json:"password"`
	IsChecklist    bool                  `json:"isChecklist"`
	ChecklistItems []model.ChecklistItem `json:"checklistItems"`
}

type NoteResponse struct {
	ID             string                `json:"id"`
	Title           string                `json:"title"`
	Content         string                `json:"content"`
	Color          model.NoteColor       `json:"color"`
	Pinned         bool                  `json:"pinned"`
	Category       string                `json:"category"`
	IsLocked        bool                  `json:"isLocked"`
	IsChecklist    bool                  `json:"isChecklist"`
	ChecklistItems []model.ChecklistItem `json:"checklistItems"`
	CreatedAt      time.Time             `json:"createdAt"`
	UpdatedAt       time.Time             `json:"updatedAt"`
}
