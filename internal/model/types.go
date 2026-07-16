package model

import "time"

type NoteColor string

const (
	ColorYellow NoteColor = "yellow"
	ColorBlue   NoteColor = "blue"
	ColorGreen  NoteColor = "green"
	ColorPink   NoteColor = "pink"
	ColorOrange NoteColor = "orange"
)

type StickyNote struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	Color     NoteColor `json:"color" bson:"color"`
	Pinned    bool      `json:"pinned" bson:"pinned"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Status         string          `json:"status"  bson:"status"`
	Category        string          `json:"category"  bson:"category"`
	IsLocked        bool            `json:"isLocked"  bson:"isLocked"`
	Password       string          `json:"password,omitempty"  bson:"password,omitempty"`
	IsChecklist     bool            `json:"isChecklist"  bson:"isChecklist"`
	ChecklistItems []ChecklistItem `json:"checklistItems" bson:"checklistItems"`
}

type NoteFormData struct {
	Title   string    `json:"title" bson:"title"`
	Content string    `json:"content" bson:"content"`
	Color   NoteColor `json:"color" bson:"color"`
}

type ChecklistItem struct {
	Text    string `json:"text" bson:"text"`
	IsDone   bool   `json:"isDone" bson:"isDone"`
}
