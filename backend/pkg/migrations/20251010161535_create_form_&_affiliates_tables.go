package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateFormsTable, downCreateFormsTable)
}

func upCreateFormsTable(ctx context.Context, tx *sql.Tx) error {
	//---- create forms table
	create_forms_table := `CREATE TABLE forms (
		id SERIAL PRIMARY KEY,
		uuid UUID DEFAULT uuid_generate_v7() NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		form_data JSONB NOT NULL,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	)`
	_, err := tx.ExecContext(ctx, create_forms_table)
	if err != nil {
		return err
	}

	//create forms table indexes
	create_forms_uuid_index := `CREATE INDEX IF NOT EXISTS idx_forms_uuid ON forms(uuid)`
	_, err = tx.ExecContext(ctx, create_forms_uuid_index)
	if err != nil {
		return err
	}

	//GIN index on form_data JSONB for better query performance
	create_form_data_index := `CREATE INDEX IF NOT EXISTS idx_forms_data ON forms USING GIN (form_data)`
	_, err = tx.ExecContext(ctx, create_form_data_index)
	if err != nil {
		return err
	}
	//---- end

	//---- create affiliates table
	create_affiliates_table := `CREATE TABLE affiliates (
		id SERIAL PRIMARY KEY,
		uuid UUID DEFAULT uuid_generate_v7() NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255),
		company VARCHAR(255),
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	)`
	_, err = tx.ExecContext(ctx, create_affiliates_table)
	if err != nil {
		return err
	}

	//create affiliates table indexes
	create_affiliates_uuid_index := `CREATE INDEX IF NOT EXISTS idx_affiliates_uuid ON affiliates(uuid)`
	_, err = tx.ExecContext(ctx, create_affiliates_uuid_index)
	if err != nil {
		return err
	}

	//---- end

	//---- create junction table for many-to-many relationship
	create_form_affiliates_table := `CREATE TABLE form_affiliates (
		form_id INTEGER NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
		affiliate_id INTEGER NOT NULL REFERENCES affiliates(id) ON DELETE CASCADE,
		added_at TIMESTAMP DEFAULT now(),
		PRIMARY KEY (form_id, affiliate_id)
	)`
	_, err = tx.ExecContext(ctx, create_form_affiliates_table)
	if err != nil {
		return err
	}

	//create junction table indexes
	create_form_affiliates_form_index := `CREATE INDEX IF NOT EXISTS idx_form_affiliates_form_id ON form_affiliates(form_id)`
	_, err = tx.ExecContext(ctx, create_form_affiliates_form_index)
	if err != nil {
		return err
	}

	create_form_affiliates_affiliate_index := `CREATE INDEX IF NOT EXISTS idx_form_affiliates_affiliate_id ON form_affiliates(affiliate_id)`
	_, err = tx.ExecContext(ctx, create_form_affiliates_affiliate_index)
	if err != nil {
		return err
	}

	//---- end

	return nil
}

func downCreateFormsTable(ctx context.Context, tx *sql.Tx) error {
	//drop tables in reverse order (junction table first due to foreign keys)
	drop_form_affiliates := `DROP TABLE IF EXISTS form_affiliates`
	_, err := tx.ExecContext(ctx, drop_form_affiliates)
	if err != nil {
		return err
	}

	drop_affiliates := `DROP TABLE IF EXISTS affiliates`
	_, err = tx.ExecContext(ctx, drop_affiliates)
	if err != nil {
		return err
	}

	drop_forms := `DROP TABLE IF EXISTS forms`
	_, err = tx.ExecContext(ctx, drop_forms)
	if err != nil {
		return err
	}

	return nil
}
