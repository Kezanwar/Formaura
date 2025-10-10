package form_repo

import (
	"encoding/json"
	"time"
)

type Model struct {
	ID          int             `db:"id"`
	UUID        string          `db:"uuid"`
	Name        string          `db:"name"`
	Description *string         `db:"description"`
	FormData    json.RawMessage `db:"form_data"` // Use json.RawMessage for JSONB
	CreatedAt   time.Time       `db:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at"`
}

// Helper method to unmarshal FormData into a specific struct
func (m *Model) UnmarshalFormData(v interface{}) error {
	return json.Unmarshal(m.FormData, v)
}
