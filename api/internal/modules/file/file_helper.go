package file

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/s3"
	"github.com/google/uuid"
)

func getFileName(fileName string) string {
	nameSplitArray := strings.Split(fileName, ".")
	fileSuffix := nameSplitArray[len(nameSplitArray)-1]
	uniqueFileName := uuid.New().String()
	completeFileName := uniqueFileName + "." + fileSuffix
	return completeFileName
}

func UploadFileAndSaveToDb(file *multipart.FileHeader, currentUser models.User) error {
	s3 := s3.GetS3()
	fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
	// => "tutorial.pdf" 360641 "application/pdf"

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	err := s3.UploadFile(ctx, getFileName(file.Filename), *file)

	if err != nil {
		return err
	}

	return nil
}
