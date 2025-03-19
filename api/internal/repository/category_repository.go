package repository

import (
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	*BaseRepository[models.Category]
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		BaseRepository: NewBaseRepository[models.Category](),
	}
}

// FindByName retrieves a category by its name
func (r *CategoryRepository) FindByName(name string, option Option) (*models.Category, error) {
	db := r.getDB(option)
	var category models.Category
	err := db.Where("name = ?", name).First(&category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

// FindByColor retrieves categories by their color
func (r *CategoryRepository) FindByColor(color string, option Option) ([]models.Category, error) {
	db := r.getDB(option)
	var categories []models.Category
	err := db.Where("color = ?", color).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
