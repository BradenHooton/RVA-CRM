package core

import (
	"time"
	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Note struct {
	BaseModel
	Note string
	NoteType string
	NoteAuthor string
	NoteDate time.Time
	NoteStatus string
	NotePriority string
	NoteCategory string
	NoteTags []string
	NoteMetadata map[string]interface{}
}
