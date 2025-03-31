package repository

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"gorm.io/gorm"
)

type FileRepository struct {
	*BaseRepository[models.File]
}

func NewFileRepository() *FileRepository {
	return &FileRepository{
		BaseRepository: NewBaseRepository[models.File](),
	}
}

func GetFileValidSortColumn() map[string]string {
	return map[string]string{
		"id":        "id",
		"createdAt": "created_at",
	}
}

func (r *FileRepository) FindWithPagination(option Option) (*PaginationResult[models.File], error) {
	searchFields := []string{"file_name", "path_name", "alt_text"}
	return r.BaseRepository.FindWithPagination(option, searchFields, GetFileValidSortColumn())
}

func (r *FileRepository) FindByUserProfileId(userId string, option Option) (*models.File, error) {
	var file *models.File
	err := r.getDB(option).Where(&models.File{UserProfileID: &userId}).First(file).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return file, err
}
