package drive_service

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"log"
	"mime/multipart"
	"time"
)

func UploadFile(driveService *drive.Service, file multipart.File, object string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	f := &drive.File{
		Name: object,
	}

	// Create and upload the file
	_, err := driveService.Files.
		Create(f).
		Media(file, googleapi.ContentType("application/octet-stream")).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()

	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
