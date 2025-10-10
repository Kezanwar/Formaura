package form_repo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, name, description string, formData interface{}) (*Model, error)
	GetByUUID(ctx context.Context, uuid string) (*Model, error)
	GetByID(ctx context.Context, id int) (*Model, error)
	FetchAll(ctx context.Context) ([]*Model, error)
	Update(ctx context.Context, uuid, name, description string, formData interface{}) (*Model, error)
	Delete(ctx context.Context, uuid string) error
}

type FormRepository struct {
	db *pgxpool.Pool
}

func NewFormRepo(db *pgxpool.Pool) *FormRepository {
	return &FormRepository{db: db}
}

func (r *FormRepository) Create(ctx context.Context, name, description string, formData interface{}) (*Model, error) {
	now := time.Now()

	// Marshal formData to JSON
	formDataJSON, err := json.Marshal(formData)
	if err != nil {
		return nil, fmt.Errorf("form.Create marshal: %w", err)
	}

	query := `
		INSERT INTO forms (name, description, form_data, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`

	var form Model

	err = pgxscan.Get(ctx, r.db, &form, query, name, description, formDataJSON, now, now)
	if err != nil {
		return nil, fmt.Errorf("form.Create query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) GetByUUID(ctx context.Context, uuid string) (*Model, error) {
	var form Model

	query := `SELECT * FROM forms WHERE uuid=$1`

	err := pgxscan.Get(ctx, r.db, &form, query, uuid)
	if err != nil {
		return nil, fmt.Errorf("form.GetByUUID query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) GetByID(ctx context.Context, id int) (*Model, error) {
	var form Model

	query := `SELECT * FROM forms WHERE id=$1`

	err := pgxscan.Get(ctx, r.db, &form, query, id)
	if err != nil {
		return nil, fmt.Errorf("form.GetByID query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) FetchAll(ctx context.Context) ([]*Model, error) {
	var forms []*Model

	query := `SELECT * FROM forms ORDER BY created_at DESC`

	err := pgxscan.Select(ctx, r.db, &forms, query)
	if err != nil {
		return nil, fmt.Errorf("form.FetchAll query: %w", err)
	}

	return forms, nil
}

func (r *FormRepository) Update(ctx context.Context, uuid, name, description string, formData interface{}) (*Model, error) {
	now := time.Now()

	// Marshal formData to JSON
	formDataJSON, err := json.Marshal(formData)
	if err != nil {
		return nil, fmt.Errorf("form.Update marshal: %w", err)
	}

	query := `
		UPDATE forms 
		SET name=$1, description=$2, form_data=$3, updated_at=$4 
		WHERE uuid=$5
		RETURNING *
	`

	var form Model

	err = pgxscan.Get(ctx, r.db, &form, query, name, description, formDataJSON, now, uuid)
	if err != nil {
		return nil, fmt.Errorf("form.Update query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) Delete(ctx context.Context, uuid string) error {
	query := `DELETE FROM forms WHERE uuid=$1`

	_, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("form.Delete: %w", err)
	}

	return nil
}
