package types

import "github.com/chiragthapa777/expense-tracker-api/internal/constant"

type ResponsePaginationMeta struct {
	Total       int64 `json:"total,omitempty"`
	Limit       int   `json:"limit,omitempty"`
	CurrentPage int   `json:"currentPage,omitempty"`
	TotalPages  int   `json:"totalPages,omitempty"`
}

type ResponseMetaData struct {
	PaginationMetaData *ResponsePaginationMeta `json:"pagination,omitempty"`
}

type Response struct {
	Success  bool                  `json:"success"`
	Data     any                   `json:"data,omitempty"`
	Error    string                `json:"error,omitempty"`
	Code     constant.ResponseCode `json:"code,omitempty"`
	MetaData *ResponseMetaData     `json:"metadata,omitempty"`
}

type ErrorResponseOption struct {
	Error  error
	Status int
	Code   constant.ResponseCode
}

type ResponseOption struct {
	Data     any
	MetaData *ResponseMetaData
	Status   int
	Code     constant.ResponseCode
}
