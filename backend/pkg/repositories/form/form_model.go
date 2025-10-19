package form_repo

import (
	"encoding/json"
	"time"
)

type FormModel struct {
	ID          int             `json:"-" db:"id"`
	UUID        string          `json:"uuid" db:"uuid"`
	UserID      int             `json:"-" db:"user_id"`
	Name        string          `json:"name" db:"name"`
	Description *string         `json:"description" db:"description"`
	FormData    json.RawMessage `json:"form_data,omitempty" db:"form_data"` // Use json.RawMessage for JSONB
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" db:"updated_at"`
}

// Helper method to unmarshal FormData into a specific struct
func (m *FormModel) UnmarshalFormData(v interface{}) error {
	return json.Unmarshal(m.FormData, v)
}
