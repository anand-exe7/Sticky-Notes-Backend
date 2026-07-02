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
	Status     string 	`json:"status"  bson:"status"`
}

type NoteFormData struct {
	Title   string    `json:"title" bson:"title"`
	Content string    `json:"content" bson:"content"`
	Color   NoteColor `json:"color" bson:"color"`
}

