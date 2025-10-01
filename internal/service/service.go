// service/file.go
package service

import (
	"context"
	"mime/multipart"
	"time"
)

// CreateFileRequest represents the data needed to create a file.
type CreateFileRequest struct {
	File       *multipart.FileHeader
	DeviceID   string
	Location   string
	UploadedAt time.Time
}

// File represents a file entity that gets stored in the database.
type File struct {
	ID         string    `json:"id"`
	Path       string    `json:"path"`
	DeviceID   string    `json:"deviceId"`
	Location   string    `json:"location"`
	UploadedAt time.Time `json:"uploadedAt"`
}

// Contracts (interfaces).
type FileService interface {
	CreateFile(ctx context.Context, req CreateFileRequest) (File, error)
	GetFile(ctx context.Context, id string) (File, error)
}

type FileRepository interface {
	Insert(ctx context.Context, f File) error
	GetByID(ctx context.Context, id string) (File, error)
}

type FileStorer interface {
	Save(file *multipart.FileHeader, filename string) (string, error)
}
