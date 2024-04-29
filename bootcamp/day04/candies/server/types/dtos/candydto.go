package dtos

type BuyCandyRequestDTO struct {
	Money      int    `json:"money" binding:"required"`
	CandyType  string `json:"candyType" binding:"required"`
	CandyCount int    `json:"candyCount" binding:"required"`
}

type BuyCandySuccessResponseDTO struct {
	Thanks string `json:"thanks"`
	Change int    `json:"change"`
}

type BuyCandyBadRequestResponseDTO struct {
	Error string `json:"error"`
}
