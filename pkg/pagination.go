package pkg

func ParsePaginationInfo(total int64, params FindRequest) *PaginationResponseMeta {
	return &PaginationResponseMeta{
		CurrentPage: params.Pagination.Page,
		PerPage:     params.Pagination.PerPage,
		TotalPages:  (total + params.Pagination.PerPage - 1) / params.Pagination.PerPage,
		TotalItems:  total,
	}
}
