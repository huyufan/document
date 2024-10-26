package main

import (
	"bytes"
	"context"
	"exec/go/exec/gostorages"
	"io"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"
)

type CustomAPIHTTPClient interface {
	Do(*http.Request) (*http.Request, error)
}

type Storage struct {
	bucket   string
	s3       *s3.Client
	uploader *manager.Uploader
}

func NewStorage(cfg Config) {

}

type Config struct {
	AccessKeyID       string
	Bucket            string
	Endpoint          string
	Region            string
	SecretAccessKey   string
	UploadConcurrency *int64
	CustomHTTPClient  CustomAPIHTTPClient
}

func (s *Storage) Save(ctx context.Context, content io.Reader, path string) error {
	input := &s3.PutObjectInput{
		ACL:    types.ObjectCannedACLPublicRead,
		Body:   content,
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}
	contentype := mime.TypeByExtension(filepath.Ext(path))

	if contentype == "" {
		data := make([]byte, 512)
		n, err := content.Read(data)
		if err != nil {
			return err
		}
		contentype = http.DetectContentType(data)
		input.Body = io.MultiReader(bytes.NewReader(data[:n]), content)
	}

	if contentype != "" {
		input.ContentType = aws.String(contentype)
	}
	_, err := s.uploader.Upload(ctx, input)
	return errors.WithStack(err)
}

func (s *Storage) Stat(ctx context.Context, path string) (*gostorages.Stat, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}
	out, err := s.s3.HeadObject(ctx, input)
	var notfounderr *types.NotFound
	if errors.As(err, &notfounderr) {
		return nil, gostorages.ErrNotExist
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	return &gostorages.Stat{
		ModifieldTime: *out.LastModified,
		Size:          *out.ContentLength,
	}, nil
}

func (s *Storage) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	input := &s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    aws.String(path),
	}
	out, err := s.s3.GetObject(ctx, input)
	var notsuckkeyerr *types.NoSuchKey
	if errors.As(err, &notsuckkeyerr) {
		return nil, gostorages.ErrNotExist
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	return out.Body, nil

}

func (s *Storage) Delete(ctx context.Context, path string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}
	_, err := s.s3.DeleteObject(ctx, input)
	return errors.WithStack(err)
}

func (s *Storage) OpenWithStat(ctx context.Context, path string) (io.ReadCloser, *gostorages.Stat, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}
	out, err := s.s3.GetObject(ctx, input)
	var notsuckkeyerr *types.NoSuchKey
	if errors.As(err, notsuckkeyerr) {
		return nil, nil, errors.Wrapf(gostorages.ErrNotExist,
			"%s does not exist in bucket %s, code: %s", path, s.bucket, notsuckkeyerr.Error())
	} else if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return out.Body, &gostorages.Stat{
		ModifieldTime: *out.LastModified,
		Size:          *out.ContentLength,
	}, nil
}
