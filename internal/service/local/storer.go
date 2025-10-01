package local

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

const uploadDir = "./uploads"

type Storer struct {
	StoragePath string
}

func NewLocalStorer() *Storer {
	return &Storer{
		StoragePath: uploadDir,
	}
}

// Save saves the uploaded file to the local filesystem and returns the file path.
func (s *Storer) Save(fileData *multipart.FileHeader, fileName string) (string, error) {
	path := getStoragePath()

	src, err := fileData.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Ensure the directory exists
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return "", err
	}

	// Create the full file path
	path += fileName

	// Create the destination file
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return path, nil
}

// getStoragePath generates a destination path for the uploaded file based on the current date and a UUID.
// The path format is: ./uploads/YYYY/MM/DD/HH/mm/.
func getStoragePath() string {
	year, month, day := time.Now().Date()
	hour, min, _ := time.Now().Clock()

	return uploadDir + "/" + strconv.Itoa(year) + "/" + month.String() + "/" + strconv.Itoa(day) + "/" + strconv.Itoa(hour) + "/" + strconv.Itoa(min) + "/"
}
