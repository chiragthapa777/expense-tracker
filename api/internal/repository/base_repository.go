package repository

import (
	"math"

	"github.com/chiragthapa777/expense-tracker-api/internal/database"
	"github.com/chiragthapa777/expense-tracker-api/internal/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	if database.DB == nil {
		panic("database.DB is not initialized")
	}
	return &BaseRepository[T]{db: database.DB}
}

// getDB returns the transaction if provided, otherwise the default DB.
func (r *BaseRepository[T]) getDB(option Option) *gorm.DB {
	if option.Tx != nil {
		return option.Tx
	}
	return r.db
}

func (r *BaseRepository[T]) Create(entity *T, option Option) error {
	db := r.getDB(option)
	// Check if the entity has an ID field and set it if empty
	if idSetter, ok := any(entity).(interface {
		GetID() string
		SetID(string)
	}); ok && idSetter.GetID() == "" {
		idSetter.SetID(uuid.New().String())
	}
	return db.Create(entity).Error
}

func (r *BaseRepository[T]) FindByID(id string, option Option) (*T, error) {
	db := r.getDB(option)
	var entity T
	err := db.Where("id = ?", id).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Or a custom error
		}
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Update(entity *T, option Option) error {
	db := r.getDB(option)
	return db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(id string, option Option) error {
	db := r.getDB(option)
	var entity T
	return db.Where("id = ?", id).Delete(&entity).Error
}

func (r *BaseRepository[T]) Find(option Option) ([]T, error) {
	db := r.getDB(option)
	var entities []T
	if err := db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepository[T]) GetQueryBuilder(option Option) (tx *gorm.DB) {
	db := r.getDB(option)
	return db.Model(new(T))
}

// FindWithPagination retrieves paginated records with optional search and sorting.
// The searchFields and validSortColumns parameters allow customization per model.
func (r *BaseRepository[T]) FindWithPagination(option Option, searchFields []string, validSortColumns map[string]string) (*PaginationResult[T], error) {
	db := r.getDB(option)

	queryBuilder := db.Model(new(T))
	if option.QueryBuilder != nil {
		queryBuilder = option.QueryBuilder
	}

	if option.PaginationDto.Search != "" {
		searchString := "%" + option.PaginationDto.Search + "%"
		for i, field := range searchFields {
			if i == 0 {
				queryBuilder.Where(field+" ILIKE ?", searchString)
			} else {
				queryBuilder.Or(field+" ILIKE ?", searchString)
			}
		}
	}

	var total int64
	if err := queryBuilder.Count(&total).Error; err != nil {
		return nil, err
	}

	limit, offset := GetLimitAndOffSet(*option.PaginationDto)
	order := GetOrderString(*option.PaginationDto, validSortColumns)
	var entities []T = make([]T, 0, limit)
	if err := queryBuilder.Limit(limit).Offset(offset).Order(order).Find(&entities).Error; err != nil {
		return nil, err
	}

	return &PaginationResult[T]{
		MetaData: types.ResponsePaginationMeta{
			Total:       total,
			Limit:       limit,
			CurrentPage: PageFromQueryString(option.PaginationDto.Page),
			TotalPages:  int(math.Ceil(float64(total) / float64(limit))),
		},
		Data: entities,
	}, nil
}
