package repository

import (
	"math"

	"github.com/chiragthapa777/expense-tracker-api/internal/database"
	"github.com/chiragthapa777/expense-tracker-api/internal/models"
	"github.com/chiragthapa777/expense-tracker-api/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func GetUserValidSortColumn() map[string]string {
	return map[string]string{
		"id":        "id",
		"email":     "email",
		"firstName": "first_name",
		"lastName":  "last_name",
	}
}

func NewUserRepository() *UserRepository {
	if database.DB == nil {
		panic("database.DB is not initialized")
	}
	return &UserRepository{db: database.DB}
}

func (r *UserRepository) Create(user *models.User, option Option) error {
	db := r.db
	if option.Tx != nil {
		db = option.Tx
	}
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	return db.Create(user).Error
}

// find user with id
func (r *UserRepository) FindByID(id string, option Option) (*models.User, error) {
	db := r.db
	if option.Tx != nil {
		db = option.Tx
	}
	var user models.User
	err := db.Where(`"users"."id" = ?`, id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Or a custom error
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User, option Option) error {
	db := r.db
	if option.Tx != nil {
		db = option.Tx
	}
	return db.Save(user).Error
}

func (r *UserRepository) Delete(id string, option Option) error {
	db := r.db
	if option.Tx != nil {
		db = option.Tx
	}
	return db.Where("id = ?", id).Delete(&models.User{}).Error
}

func (r *UserRepository) Find(option Option) ([]models.User, error) {
	db := r.db
	if option.Tx != nil {
		db = option.Tx
	}
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindWithPagination(option Option) (*PaginationResult[models.User], error) {
	db := r.db
	if option.Tx != nil {
		db = option.Tx
	}

	var total int64
	limit, offset := GetLimitAndOffSet(*option.PaginationDto)
	order := GetOrderString(*option.PaginationDto, GetUserValidSortColumn())

	var users []models.User = make([]models.User, 0, limit)

	queryBuilder := db.Model(&models.User{})

	if option.PaginationDto.Search != "" {
		searchString := "%" + option.PaginationDto.Search + "%"
		queryBuilder.Where(" email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ? ", searchString, searchString, searchString)
	}

	if err := queryBuilder.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := queryBuilder.Limit(limit).Offset(offset).Order(order).Find(&users).Error; err != nil {
		return nil, err
	}

	return &PaginationResult[models.User]{
		MetaData: utils.ResponsePaginationMeta{
			Total:       total,
			Limit:       limit,
			CurrentPage: PageFromQueryString(option.PaginationDto.Page),
			TotalPages:  int(math.Ceil(float64(total) / float64(limit))),
		},
		Data: users,
	}, nil
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

func (r *UserRepository) UpdatePassword(userId string, newPassword string, option Option) error {
	db := r.db
	if option.Tx != nil {
		db = option.Tx
	}

	updateData := map[string]any{
		"password":                newPassword,
		"is_password_set_by_user": true,
	}

	return db.Model(&models.User{}).Where("id = ?", userId).Updates(updateData).Error
}
