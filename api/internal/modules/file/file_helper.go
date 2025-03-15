package file

import (
	"context"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/repository"
	"github.com/chiragthapa777/expense-tracker-api/internal/s3"
	"github.com/google/uuid"
)

func GetFileMimeType(file *multipart.FileHeader) (string, error) {

	// Open file to detect content type
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	buffer := make([]byte, 512)
	n, err := f.Read(buffer)
	if err != nil {
		return "", err
	}

	// Detect content type
	contentType := http.DetectContentType(buffer[:n])

	// Reset file reader for downstream handlers
	_, err = f.Seek(0, 0)
	if err != nil {
		return "", err
	}

	return contentType, nil
}

func getFileName(fileName string) string {
	nameSplitArray := strings.Split(fileName, ".")
	fileSuffix := nameSplitArray[len(nameSplitArray)-1]
	uniqueFileName := uuid.New().String()
	completeFileName := uniqueFileName + "." + fileSuffix
	return completeFileName
}

func UploadFileAndSaveToDb(file *multipart.FileHeader, currentUser models.User) (*models.File, error) {
	s3 := s3.GetS3()
	fileRepository := repository.NewFileRepository()

	mimeType, err := GetFileMimeType(file)
	if err != nil {
		return nil, err
	}

	fileName := getFileName(file.Filename)
	newFile := models.File{
		FileName:  fileName,
		MimeType:  mimeType,
		PathName:  currentUser.ID,
		IsPrivate: true,
		Variants: models.JSONB{
			"thumbnail": models.FileVariantData{},
			"small":     models.FileVariantData{},
			"medium":    models.FileVariantData{},
			"large":     models.FileVariantData{},
		},
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	objectKey := newFile.PathName + "/" + newFile.FileName

	if err = s3.UploadFile(ctx, objectKey, &newFile.MimeType, *file); err != nil {
		return nil, err
	}

	if err = fileRepository.Create(&newFile, repository.Option{}); err != nil {
		return nil, err
	}

	return &newFile, err
}
