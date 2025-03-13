package repository

import (
	"fmt"
	"strconv"

	"github.com/chiragthapa777/expense-tracker-api/internal/dto"
	"github.com/chiragthapa777/expense-tracker-api/internal/utils"
)

func GetLimitAndOffSet(paginationOption dto.PaginationQueryDto) (int, int) {
	offset := 0
	limit := LimitFromQueryString(paginationOption.Limit)
	page := PageFromQueryString(paginationOption.Page)

	offset = (page - 1) * limit

	return limit, offset
}

func PageFromQueryString(pageStr string) int {
	page := 1
	pageFromQuery, _ := strconv.Atoi(pageStr)
	if pageFromQuery > 0 {
		page = pageFromQuery
	}
	return page
}

func LimitFromQueryString(limitStr string) int {
	limit := 20
	limitFromQuery, _ := strconv.Atoi(limitStr)
	if limitFromQuery > 0 {
		limit = utils.MathMinInt(100, limitFromQuery)
	}
	return limit
}

// GetOrderString builds a SQL ORDER BY clause from pagination options.
//
// Uses validColumn to map SortBy to a safe column name (defaults to "created_at")
// and sets sortOrder to "asc" if specified, else "desc".
//
// Args:
//   - paginationOption: Contains SortBy and SortOrder.
//   - validColumn: Maps user sort fields to database columns.
//
// Returns:
//
//	A string like "<column> <order>" (e.g., "created_at desc").
//
// Example:
//
//	opts := dto.PaginationQueryDto{SortBy: "date", SortOrder: "asc"}
//	validColumn := map[string]string{"date": "created_at", "name": "first_name"}
//	result := GetOrderString(validColumn, opts) // Returns "created_at asc"
func GetOrderString(paginationOption dto.PaginationQueryDto, validColumn map[string]string) string {
	sortOrder, sortBy := "desc", "created_at"
	if paginationOption.SortOrder != nil && *paginationOption.SortOrder == "asc" {
		sortOrder = "asc"
	}
	if paginationOption.SortBy != "" {
		val, ok := validColumn[paginationOption.SortBy]
		if ok {
			sortBy = val
		}
	}
	return fmt.Sprintf("%s %s", sortBy, sortOrder)
}
