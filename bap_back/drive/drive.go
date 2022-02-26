package drive

import (
	"context"
	"net/http"

	driveAPI "google.golang.org/api/drive/v3"
)

type Drive interface {
	Construct(ctx context.Context) error
	Export(fileId, mimeType string, ctx context.Context) (*http.Response, error)
}

type DriveImpl struct {
	FilesSrv *driveAPI.FilesService
}

func NewDrive() Drive {
	return &DriveImpl{FilesSrv: nil}
}

func (drive *DriveImpl) Construct(ctx context.Context) error {
	srv, err := driveAPI.NewService(ctx)
	if err != nil {
		return err
	}
	filesSrv := driveAPI.NewFilesService(srv)

	drive.FilesSrv = filesSrv
	return nil
}

func (drive *DriveImpl) Export(fileId, mimeType string, ctx context.Context) (*http.Response, error) {
	return drive.FilesSrv.Export(fileId, mimeType).Context(ctx).Download()
}
