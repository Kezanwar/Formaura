package form_repo

import (
	"encoding/json"
	"time"
)

type ListingModel struct {
	UUID            string          `db:"uuid" json:"uuid"`
	Name            string          `db:"name" json:"name"`
	Description     *string         `db:"description" json:"description"`
	CreatedAt       time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time       `db:"updated_at" json:"updated_at"`
	Affiliates      json.RawMessage `db:"affiliates" json:"affiliates"`
	SubmissionCount int             `db:"submission_count" json:"submission_count"`
}

func (m *ListingModel) GetAffiliates() ([]AffiliateInfo, error) {
	var affiliates []AffiliateInfo
	err := json.Unmarshal(m.Affiliates, &affiliates)
	return affiliates, err
}

type AffiliateInfo struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
