package types

import "server/types/models"

type PaginatedRestaurants struct {
	Items  []models.Restaurant `json:"items"`
	Paging PagingOutgoing      `json:"paging"`
}
