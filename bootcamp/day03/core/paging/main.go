package paging

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"server/types"
)

var (
	DEFAULT_PAGE_SIZE   = 10
	PAGE_SIZE_KEY       = "pageSize"
	DEFAULT_PAGE_NUMBER = 1
	PAGE_NUMBER_KEY     = "pageNumber"
	PAGIONATION_KEY     = "pagination"
)

func GetPaginationFromQueryParams(c *gin.Context) (types.PagingIncoming, error) {
	res := types.PagingIncoming{}
	pageNumQuery := c.DefaultQuery(PAGE_NUMBER_KEY, fmt.Sprint(DEFAULT_PAGE_NUMBER))
	pageSizeQuery := c.DefaultQuery(PAGE_SIZE_KEY, fmt.Sprint(DEFAULT_PAGE_SIZE))
	pageSize, sizeErr := strconv.Atoi(pageSizeQuery)
	pageNumber, numberErr := strconv.Atoi(pageNumQuery)
	if sizeErr != nil || numberErr != nil {
		return res, errors.New(
			fmt.Sprintf(
				"pagination query parameters must be numbers. Error in one of %s, %s",
				PAGE_SIZE_KEY,
				PAGE_NUMBER_KEY,
			),
		)
	}

	res.PageNumber = pageNumber
	res.PageSize = pageSize
	if pageNumber < 1 || pageSize < 1 {
		return res, errors.New("pageSize and pageNumber can't be smaller than 1")
	}
	return res, nil
}

func GetPaginationFromJsonBody() (types.PagingIncoming, error) {
	// TODO
	return types.PagingIncoming{}, nil
}
