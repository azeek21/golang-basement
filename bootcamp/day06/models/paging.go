package models

import (
	"fmt"
	"math"
)

const PAGING_MODEL_NAME = "paging"

var PAGING_PAGE_NAME = "page"
var PAGING_ITEMS_PER_PAGE_NAME = "itemsPerPage"
var PAGING_TOTAL_MATCHES_COUNT_NAME = "total"

const PAGING_DEFAULT_ITEMS_PER_PAGE = 5
const PAGING_DEFAULT_PAGE = 1
const DEFAULT_PAGING_QUERY_PARAM = "?page=1&itemsPerPage=10"

type PagingIncoming struct {
	Page         int `form:"page" binding:"required"`
	ItemsPerPage int `form:"itemsPerPage" binding:"required"`
}

type PagingOutGoing struct {
	Page         int
	ItemsPerPage int
	Total        int
}

func (p PagingIncoming) TransformToOutgoing(total int) *PagingOutGoing {
	return &PagingOutGoing{
		Page:         p.Page,
		ItemsPerPage: p.ItemsPerPage,
		Total:        total,
	}
}

func NewDefaultPagingIncoming() *PagingIncoming {
	return &PagingIncoming{
		Page:         PAGING_DEFAULT_PAGE,
		ItemsPerPage: PAGING_DEFAULT_ITEMS_PER_PAGE,
	}
}

func (p PagingOutGoing) HasNextPage() bool {
	pageCount := math.Ceil(float64(p.Total) / float64(p.ItemsPerPage))
	return p.Page < int(pageCount)
}

func (p PagingOutGoing) HasPreviousPage() bool {
	return p.Page > 1
}

func GeneratePagingQuery(page, itemPerPage int) string {
	return fmt.Sprintf("?%s=%d&%s=%d", PAGING_PAGE_NAME, page, PAGING_ITEMS_PER_PAGE_NAME, itemPerPage)
}

func (p PagingOutGoing) GetCurrentPageUrlParamsAsString() string {
	// ?page=3&ItemsPerPage=10
	return GeneratePagingQuery(p.Page, p.ItemsPerPage)
}

func (p PagingOutGoing) GetNextPageUrlParamsAsString() string {
	if p.HasNextPage() {
		return GeneratePagingQuery(p.Page+1, p.ItemsPerPage)
	}
	return p.GetCurrentPageUrlParamsAsString()
}

func (p PagingOutGoing) GetPreviousPageUrlParamsAsString() string {
	if p.HasPreviousPage() {
		return GeneratePagingQuery(p.Page-1, p.ItemsPerPage)
	}
	return p.GetCurrentPageUrlParamsAsString()
}
