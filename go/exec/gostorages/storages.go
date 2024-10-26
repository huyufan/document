package gostorages

import (
	"context"
	"errors"
	"io"
	"time"
)

type Storage interface {
	Save(ctx context.Context, content io.Reader, path string) error
	Stat(ctx context.Context, path string) (*Stat, error)
	Open(ctx context.Context, path string) (io.ReadCloser, error)
	OpenWithStat(ctx context.Context, path string) (io.ReadCloser, *Stat, error)
	Delete(ctx context.Context, path string) error
}

type Stat struct {
	ModifieldTime time.Time
	Size          int64
}

var ErrNotExist = errors.New("does not exist")
