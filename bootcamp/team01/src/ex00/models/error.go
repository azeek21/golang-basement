package models

type BadRequestDTO struct {
	Error string `json:"error"`
}

func NewBadRequestResponse(err error) *BadRequestDTO {
	res := BadRequestDTO{}
	if err != nil {
		res.Error = err.Error()
		return &res
	}
	return &res
}
