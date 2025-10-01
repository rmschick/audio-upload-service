package local

import (
	"context"
	"time"

	"github.com/google/uuid"

	"personal-dev/internal/service"
)

type FileService struct {
	Repo   service.FileRepository
	Storer service.FileStorer
}

// NewFileService creates a new instance of FileService.
func NewFileService(repo service.FileRepository, storer service.FileStorer) *FileService {
	return &FileService{
		Repo:   repo,
		Storer: storer,
	}
}

// CreateFile handles the logic of saving a file and creating its database record.
func (s *FileService) CreateFile(ctx context.Context, req service.CreateFileRequest) (service.File, error) {
	filename := req.File.Filename

	path, err := s.Storer.Save(req.File, filename)
	if err != nil {
		return service.File{}, err
	}

	f := service.File{
		ID:         uuid.NewString(),
		Path:       path,
		DeviceID:   req.DeviceID,
		Location:   req.Location,
		UploadedAt: time.Now(),
	}

	if err := s.Repo.Insert(ctx, f); err != nil {
		return service.File{}, err
	}

	return f, nil
}

// GetFile retrieves a file by its ID.
func (s *FileService) GetFile(ctx context.Context, id string) (service.File, error) {
	return s.Repo.GetByID(ctx, id)
}
