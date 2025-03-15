package repository

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
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
