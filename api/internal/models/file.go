package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/chiragthapa777/expense-tracker-api/internal/s3"
	"gorm.io/gorm"
)

// JSONB type for handling JSONB in GORM

type FileVariantData struct {
	FileName  string  `json:"fileName"`
	PathName  string  `json:"pathName"`
	SignedUrl *string `json:"signedUrl"`
}

type JSONB map[string]FileVariantData

// Scan implements the Scanner interface for JSONB
func (j *JSONB) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

// Value implements the driver Valuer interface for JSONB
func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type File struct {
	BaseModel
	MimeType  string `gorm:"type:varchar;not null" json:"mimeType"`
	FileName  string `gorm:"type:varchar;unique" json:"fileName"`
	PathName  string `gorm:"type:varchar;not null" json:"pathName"`
	AltText   string `gorm:"type:text" json:"altText"`
	IsPrivate bool   `gorm:"not null;default:false" json:"isPrivate"`
	Variants  JSONB  `gorm:"type:jsonb" json:"variants"`

	SignedUrl *string `gorm:"-" json:"signedUrl,omitempty"`
}

func (f *File) SetPresignedUrl() (err error) {
	s3 := s3.GetS3()
	objectKey := f.PathName + "/" + f.FileName
	presignedUrl, err := s3.GetObject(context.TODO(), objectKey, 24*60*60)
	if err != nil {
		return err
	}
	for key, v := range f.Variants {
		if v.FileName != "" && v.PathName != "" {
			objectKey := f.PathName + "/" + f.FileName
			presignedUrl, err := s3.GetObject(context.TODO(), objectKey, 24*60*60)
			if err != nil {
				return nil
			}
			v.SignedUrl = &presignedUrl.URL
			f.Variants[key] = v
		}
	}
	f.SignedUrl = &presignedUrl.URL
	return nil
}

func (f *File) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("AfterFind")
	f.SetPresignedUrl()
	return nil
}

func (f *File) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate")
	f.SetPresignedUrl()
	return nil
}
