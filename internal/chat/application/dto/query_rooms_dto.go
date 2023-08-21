package dto

type QueryRoomsReqDTO struct {
	// Filters
	UserIn    []string
	UserNotIn []string

	// Pagination
	Page    int `json:"page" validate:"gte=1"`
	PerPage int `json:"per_page" validate:"gt=0"`

	// Sorting
	SortField string `json:"sort_field"`
	SortOrder string `json:"sort_order" validate:"required,oneof=asc desc"`
}
