package fs

import (
	"context"
	"exec/go/exec/gostorages"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Storage struct {
	root string
}

func NewStorage(cfg Config) *Storage {
	return &Storage{root: cfg.Root}
}

type Config struct {
	Root string
}

func (fs *Storage) abs(path string) string {
	return filepath.Join(fs.root, path)
}

func (fs *Storage) Save(ctx context.Context, content io.Reader, path string) error {
	abs := fs.abs(path)
	if err := os.MkdirAll(filepath.Dir(abs), 0755); err != nil {
		return errors.WithStack(err)
	}
	w, err := os.Create(abs)
	if err != nil {
		return errors.WithStack(err)
	}
	defer w.Close()
	if _, err := io.Copy(w, content); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (fs *Storage) Stat(ctx context.Context, path string) (*gostorages.Stat, error) {
	fi, err := os.Stat(fs.abs(path))
	if os.IsNotExist(err) {
		return nil, gostorages.ErrNotExist
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	return &gostorages.Stat{
		ModifieldTime: fi.ModTime(),
		Size:          fi.Size(),
	}, nil
}

func (fs *Storage) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	f, err := os.Open(fs.abs(path))
	if os.IsNotExist(err) {
		return nil, gostorages.ErrNotExist
	}
	return f, errors.WithStack(err)
}

func (fs *Storage) Delete(ctx context.Context, path string) error {
	return os.Remove(fs.abs(path))
}

func (fs *Storage) OpenWithStat(ctx context.Context, path string) (io.ReadCloser, *gostorages.Stat, error) {
	f, err := os.Open(fs.abs(path))
	if os.IsNotExist(err) {
		return nil, nil, gostorages.ErrNotExist
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return f, &gostorages.Stat{
		ModifieldTime: stat.ModTime(),
		Size:          stat.Size(),
	}, nil
}
