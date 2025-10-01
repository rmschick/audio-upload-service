package local

import (
	"context"
	"database/sql"
	"fmt"

	"personal-dev/internal/service"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

// Insert adds a new file record to the database.
func (r *PostgresRepository) Insert(ctx context.Context, f service.File) error {
	_, err := r.DB.ExecContext(ctx,
		"INSERT INTO uploads (id, path, device_id, location, uploaded_at) VALUES ($1, $2, $3, $4, $5)",
		f.ID, f.Path, f.DeviceID, f.Location, f.UploadedAt,
	)
	
	return err
}

// GetByID retrieves a file by its ID.
func (r *PostgresRepository) GetByID(ctx context.Context, id string) (service.File, error) {
	var f service.File
	
	row := r.DB.QueryRowContext(ctx, "SELECT id, path, device_id, location, uploaded_at FROM uploads WHERE id=$1", id)
	if err := row.Scan(&f.ID, &f.Path, &f.DeviceID, &f.Location, &f.UploadedAt); err != nil {
		return f, fmt.Errorf("query failed: %w", err)
	}

	return f, nil
}
