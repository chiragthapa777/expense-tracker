package repository

import (
	"time"

	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository[models.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[models.User](),
	}
}

// GetUserValidSortColumn returns valid sort columns for User.
func GetUserValidSortColumn() map[string]string {
	return map[string]string{
		"id":        "id",
		"email":     "email",
		"firstName": "first_name",
		"lastName":  "last_name",
		"createdAt": "created_at",
	}
}

// FindWithPagination overrides the base method to specify User-specific search fields.
func (r *UserRepository) FindWithPagination(option Option) (*PaginationResult[models.User], error) {
	searchFields := []string{"email", "first_name", "last_name"}
	return r.BaseRepository.FindWithPagination(option, searchFields, GetUserValidSortColumn())
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Or a custom error
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByIdWithJoins(id string) (*models.User, error) {
	db := r.db
	var entity models.User
	err := db.Joins("Profile").Where(&models.User{BaseModel: models.BaseModel{ID: id}}).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Or a custom error
		}
		return nil, err
	}
	return &entity, nil
}

func (r *UserRepository) UpdatePassword(userId string, newPassword string, option Option) error {
	db := r.getDB(option)

	updateData := map[string]any{
		"password":                newPassword,
		"is_password_set_by_user": true,
	}

	return db.Model(&models.User{}).Where("id = ?", userId).Updates(updateData).Error
}

func (r *UserRepository) BlockUser(userId string, option Option) error {
	db := r.getDB(option)

	updateData := map[string]any{
		"blocked_at": time.Now(),
	}

	return db.Model(&models.User{}).Where("id = ?", userId).Updates(updateData).Error
}

func (r *UserRepository) UnBlockUser(userId string, option Option) error {
	db := r.getDB(option)

	updateData := map[string]any{
		"blocked_at": nil,
	}

	return db.Model(&models.User{}).Where("id = ?", userId).Updates(updateData).Error
}
