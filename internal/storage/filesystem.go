package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileStorage interface {
	Save(ctx context.Context, path string, file io.Reader) error
	Get(ctx context.Context, path string) (io.ReadCloser, error)
	Delete(ctx context.Context, path string) error
}

type LocalFileStorage struct {
	basePath    string
	publicURL   string
	maxFileSize int64
}

type Config struct {
	BasePath    string
	PublicURL   string
	MaxFileSize int64
}

func NewLocalFileStorage(config *Config) (*LocalFileStorage, error) {

	if err := os.MkdirAll(config.BasePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	return &LocalFileStorage{
		basePath:    config.BasePath,
		publicURL:   config.PublicURL,
		maxFileSize: config.MaxFileSize,
	}, nil
}

func (fs *LocalFileStorage) Save(ctx context.Context, path string, file io.Reader) error {
	fullpath := filepath.Join(fs.basePath, filepath.Clean(path))

	dir := filepath.Dir(fullpath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	tmpPath := fullpath + ".tmp"
	f, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer func() {
		f.Close()
		if err != nil {
			os.Remove(tmpPath)
		}
	}()

	written, err := io.Copy(f, io.LimitReader(file, fs.maxFileSize))
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	if written >= fs.maxFileSize {
		return fmt.Errorf("file too large: maximum size is %d bytes", fs.maxFileSize)
	}

	if err := f.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	if err := os.Rename(tmpPath, fullpath); err != nil {
		return fmt.Errorf("failed to move file to final location: %w", err)
	}

	return nil
}

func (fs *LocalFileStorage) Get(ctx context.Context, path string) (io.ReadCloser, error) {
	fullpath := filepath.Join(fs.basePath, filepath.Clean(path))

	file, err := os.Open(fullpath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file not found: %w", err)
		}
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	return file, nil
}
