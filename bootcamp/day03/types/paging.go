package types

type PagingIncoming struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
}

type PagingOutgoing struct {
	PagingIncoming
	Total int64 `json:"total"`
}
