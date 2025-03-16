package file

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"
	"sync"

	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/modules/auth"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/response"
	"github.com/chiragthapa777/expense-tracker-api/internal/s3"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadImages(c *fiber.Ctx) error {
	currentUser, err := auth.CurrentUser(c)
	if err != nil || currentUser == nil {
		return response.SendError(c, types.ErrorResponseOption{Error: errors.New("failed to get current user")})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	var wg sync.WaitGroup
	errChan := make(chan error, len(files))
	fileChan := make(chan models.File, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader, currentUser models.User) {
			defer wg.Done()
			newFile, err := UploadFileAndSaveToDb(file, currentUser)
			if err != nil {
				errChan <- err
			} else {
				fileChan <- *newFile
			}
		}(file, *currentUser)
	}

	wg.Wait()
	close(errChan)
	close(fileChan)

	var errors []error
	for err := range errChan {
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return response.SendError(c, types.ErrorResponseOption{Error: fmt.Errorf("upload failed: %v", errors)})
	}

	var newFiles []models.File
	for file := range fileChan {
		newFiles = append(newFiles, file)
	}

	return response.Send(c, types.ResponseOption{Data: newFiles})
}

func DeleteFiles(c *fiber.Ctx) error {
	ids := c.Params("ids")
	if ids == "" {
		return response.SendError(c, types.ErrorResponseOption{Error: errors.New("file id is required"), Status: fiber.StatusBadRequest})
	}
	fileIds := strings.Split(ids, ",")
	for _, fileId := range fileIds {
		if err := uuid.Validate(fileId); err != nil {
			return response.SendError(c, types.ErrorResponseOption{Error: err, Status: fiber.StatusBadRequest})
		}
	}

	fileRepository := repository.NewFileRepository()
	s3 := s3.GetS3()

	for _, fileId := range fileIds {
		foundFile, err := fileRepository.FindByID(fileId, repository.Option{})
		if err != nil {
			return response.SendError(c, types.ErrorResponseOption{Error: err})
		}
		if foundFile == nil {
			return response.SendError(c, types.ErrorResponseOption{Error: repository.ErrRecordNotFound})
		}
		fileRepository.Delete(foundFile.ID, repository.Option{})
		key := foundFile.PathName + "/" + foundFile.FileName
		s3.DeleteObject(context.TODO(), key, "", false)
	}

	return response.Send(c, types.ResponseOption{Data: "Files deleted successfully"})
}
