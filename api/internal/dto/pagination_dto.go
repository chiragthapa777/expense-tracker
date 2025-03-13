package dto

type PaginationQueryDto struct {
	Page      string  `query:"page" validate:"numeric"`
	Limit     string  `query:"limit" validate:"numeric"`
	Search    string  `query:"search"`
	SortBy    string  `query:"sortBy"`
	SortOrder *string `query:"sortOrder" validate:"omitempty,oneof=asc desc"`
}
