package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/maksymchuk-mm/crm/pkg/schemas"
	"strconv"
)

func GeneratePaginationRequest(c *gin.Context) *schemas.Pagination {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sort := c.DefaultQuery("sort", "created_at desc")
	return &schemas.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
