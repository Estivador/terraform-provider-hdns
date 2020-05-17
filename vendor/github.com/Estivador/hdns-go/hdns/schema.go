package hdns

import (
	"github.com/Estivador/hdns-go/hdns/schema"
)

// PaginationFromSchema converts a schema.MetaPagination to a Pagination.
func PaginationFromSchema(s schema.MetaPagination) Pagination {
	return Pagination{
		Page:         s.Page,
		PerPage:      s.PerPage,
		LastPage:     s.LastPage,
		TotalEntries: s.TotalEntries,
	}
}
