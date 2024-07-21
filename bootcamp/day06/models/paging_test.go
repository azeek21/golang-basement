package models_test

import (
	"testing"

	"github.com/azeek21/blog/models"
	"github.com/stretchr/testify/assert"
)

type PagingCase struct {
	In    *models.PagingOutGoing
	Out   any
	Title string
}

func TestCurrentPageQueryParams(t *testing.T) {
	test_cases := []PagingCase{
		{
			In:    &models.PagingOutGoing{},
			Out:   "?page=0&itemsPerPage=0",
			Title: "Everything is 0 or EMPTY PAGING",
		},
		{
			In: &models.PagingOutGoing{
				Page:         1,
				ItemsPerPage: 10,
			},
			Out:   "?page=1&itemsPerPage=10",
			Title: "With correct values",
		},
		{
			In: &models.PagingOutGoing{
				Page:         12345,
				ItemsPerPage: 100,
			},
			Out:   "?page=12345&itemsPerPage=100",
			Title: "With random values",
		},
	}

	for _, testCase := range test_cases {
		t.Run(testCase.Title, func(t *testing.T) {
			assert.Equal(t, testCase.Out, testCase.In.GetCurrentPageUrlParamsAsString())
		})
	}
}

func TestPagingHasNextPage(t *testing.T) {
	test_cases := []PagingCase{
		{
			In:    &models.PagingOutGoing{},
			Out:   false,
			Title: "1. Everything is 0 or EMPTY PAGING",
		},
		{
			In: &models.PagingOutGoing{
				Page:         1,
				ItemsPerPage: 10,
				Total:        11,
			},
			Out:   true,
			Title: "2. Should have next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         1,
				ItemsPerPage: 10,
				Total:        19,
			},
			Out:   true,
			Title: "3. Should have next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         2,
				ItemsPerPage: 18,
				Total:        19,
			},
			Out:   false,
			Title: "4. Should have next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         10,
				ItemsPerPage: 10,
				Total:        101,
			},
			Out:   true,
			Title: "5. Should have next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         4,
				ItemsPerPage: 4,
				Total:        33,
			},
			Out:   true,
			Title: "6. Should have next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         8,
				ItemsPerPage: 4,
				Total:        33,
			},
			Out:   true,
			Title: "7. Should have next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         9,
				ItemsPerPage: 4,
				Total:        33,
			},
			Out:   false,
			Title: "8. Should NOT have next page, we already at last page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         11,
				ItemsPerPage: 10,
				Total:        101,
			},
			Out:   false,
			Title: "9. Should NOT have next page, we already at last page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         2,
				ItemsPerPage: 10,
				Total:        19,
			},
			Out:   false,
			Title: "10. Should NOT have next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         5,
				ItemsPerPage: 10,
				Total:        59,
			},
			Out:   true,
			Title: "11. Should have next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         6,
				ItemsPerPage: 10,
				Total:        59,
			},
			Out:   false,
			Title: "12. Should NOT have next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         6,
				ItemsPerPage: 10,
				Total:        61,
			},
			Out:   true,
			Title: "13. Should have next page",
		},
	}

	for _, testCase := range test_cases {
		t.Run(testCase.Title, func(t *testing.T) {
			assert.Equal(t, testCase.Out, testCase.In.HasNextPage())
		})
	}
}

func TestGetNextPageUrlParams(t *testing.T) {
	test_cases := []PagingCase{
		{
			In:    &models.PagingOutGoing{},
			Out:   "?page=0&itemsPerPage=0",
			Title: "1. Everything is 0 or EMPTY PAGING",
		},
		{
			In: &models.PagingOutGoing{
				Page:         1,
				ItemsPerPage: 10,
				Total:        10,
			},
			Out:   "?page=1&itemsPerPage=10",
			Title: "2. NO next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         1,
				ItemsPerPage: 10,
				Total:        11,
			},
			Out:   "?page=2&itemsPerPage=10",
			Title: "3. HAS next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         2,
				ItemsPerPage: 10,
				Total:        11,
			},
			Out:   "?page=2&itemsPerPage=10",
			Title: "3. NO next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         10,
				ItemsPerPage: 10,
				Total:        101,
			},
			Out:   "?page=11&itemsPerPage=10",
			Title: "3. HAS next page",
		},
		{
			In: &models.PagingOutGoing{
				Page:         11,
				ItemsPerPage: 10,
				Total:        101,
			},
			Out:   "?page=11&itemsPerPage=10",
			Title: "3. NO next page",
		},
	}

	for _, testCase := range test_cases {
		t.Run(testCase.Title, func(t *testing.T) {
			assert.Equal(t, testCase.Out, testCase.In.GetNextPageUrlParamsAsString())
		})
	}
}
