package drive_service

import (
	"context"
	"errors"
	"gdrive_uploader/pkg/auth"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func NewService() (*drive.Service, error) {
	ctx := context.Background()
	client, err := auth.Client()
	if err != nil {
		return nil, err
	}

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, errors.New("creating drive_service service failed")
	}

	return srv, nil
}
