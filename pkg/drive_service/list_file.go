package drive_service

import (
	"errors"
	"google.golang.org/api/drive/v3"
)

func ListFile(driveService *drive.Service) (*drive.FileList, error) {
	r, err := driveService.Files.List().
		PageSize(10).
		Fields("nextPageToken, files(id, name, size, createdTime)").
		OrderBy("name").
		Do()
	if err != nil {
		return nil, errors.New("Unable to retrieve files")
	}

	if len(r.Files) == 0 {
		return nil, errors.New("No files found.")
	}

	return r, nil
}
