package form_repo

import (
	"context"
	"encoding/json"
	"fmt"
	"formaura/pkg/util"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, userId int, name string, description *string, formData FormData) (*FormModel, error)
	GetByUUID(ctx context.Context, uuid string) (*FormModel, error)
	GetByID(ctx context.Context, id int) (*FormModel, error)
	GetBasicListingByUserID(ctx context.Context, id int) ([]*FormModel, error)
	GetDetailedListingByUserID(ctx context.Context, id int) ([]*ListingModel, error)
	UpdateFormMeta(ctx context.Context, id int, name, description string) (*FormModel, error)
	Delete(ctx context.Context, uuid string) error
}

type FormRepository struct {
	db *pgxpool.Pool
}

func NewFormRepo(db *pgxpool.Pool) *FormRepository {
	return &FormRepository{db: db}
}

func (r *FormRepository) Create(ctx context.Context, user_id int, name string, description *string, formData FormData) (*FormModel, error) {
	now := time.Now()

	// Marshal formData to JSON
	formDataJSON, err := json.Marshal(formData)
	if err != nil {
		return nil, fmt.Errorf("form.Create marshal: %w", err)
	}

	query := `
		INSERT INTO forms (user_id, name, description, form_data, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
	`

	var form FormModel

	err = pgxscan.Get(ctx, r.db, &form, query, user_id, name, description, formDataJSON, now, now)

	util.PrintStruct(form)

	if err != nil {
		return nil, fmt.Errorf("form.Create query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) GetByUUID(ctx context.Context, uuid string) (*FormModel, error) {
	var form FormModel

	query := `
	SELECT * 
	FROM forms 
	WHERE uuid=$1`

	err := pgxscan.Get(ctx, r.db, &form, query, uuid)
	if err != nil {
		return nil, fmt.Errorf("form.GetByUUID query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) GetByID(ctx context.Context, id int) (*FormModel, error) {
	var form FormModel

	query := `
	SELECT * 
	FROM forms 
	WHERE id=$1`

	err := pgxscan.Get(ctx, r.db, &form, query, id)
	if err != nil {
		return nil, fmt.Errorf("form.GetByID query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) GetBasicListingByUserID(ctx context.Context, id int) ([]*FormModel, error) {
	forms := []*FormModel{}

	query := `
	SELECT uuid, name, description, created_at, updated_at 
	FROM forms 
	WHERE user_id = $1
	ORDER BY created_at DESC`

	err := pgxscan.Select(ctx, r.db, &forms, query, id)
	if err != nil {
		return nil, fmt.Errorf("form.FetchAll query: %w", err)
	}

	return forms, nil
}

func (r *FormRepository) GetDetailedListingByUserID(ctx context.Context, id int) ([]*ListingModel, error) {
	forms := []*ListingModel{}

	query := `
	SELECT 
	f.uuid,
	f.name,
	f.description,
	f.created_at,
	f.updated_at,
	COALESCE(
		jsonb_agg(
			jsonb_build_object(
				'uuid', a.uuid,
				'first_name', a.first_name,
				'last_name', a.last_name
				)
			) FILTER (WHERE a.uuid IS NOT NULL),
			'[]'::jsonb
		) as affiliates,
		COUNT(DISTINCT fs.id) as submission_count
	FROM forms f
	LEFT JOIN form_affiliates fa ON f.id = fa.form_id
	LEFT JOIN affiliates a ON fa.affiliate_id = a.id
	LEFT JOIN form_submissions fs ON f.id = fs.form_id
	WHERE f.user_id = $1
	GROUP BY f.id, f.uuid, f.name, f.description, f.created_at, f.updated_at
	ORDER BY f.created_at DESC`

	err := pgxscan.Select(ctx, r.db, &forms, query, id)
	if err != nil {
		return nil, fmt.Errorf("form.GetDetailedListingByUserID query: %w", err)
	}

	return forms, nil
}
func (r *FormRepository) UpdateFormMeta(ctx context.Context, id int, name, description string) (*FormModel, error) {
	now := time.Now()
	fmt.Println(description)
	query := `
		UPDATE forms 
		SET name=$1, description=$2, updated_at=$3 
		WHERE id=$4
		RETURNING *
	`

	var form FormModel

	err := pgxscan.Get(ctx, r.db, &form, query, name, description, now, id)

	if err != nil {
		return nil, fmt.Errorf("form.Update query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) UpdateFormData(ctx context.Context, uuid, name, description string, formData FormData) (*FormModel, error) {
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

	var form FormModel

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
