package models

type Candy struct {
	Code      string
	Name      string `jsone:"name"`
	Price     int    `json:"price"`
	PriceUnit string `json:"priceUnit"`
}

var CANDIES = map[string]*Candy{
	"CE": {
		Code:      "CE",
		Name:      "Cool Eskimo",
		Price:     10,
		PriceUnit: "cents",
	},
	"AA": {
		Code:      "AA",
		Name:      "Apricont Aardvark",
		Price:     15,
		PriceUnit: "cents",
	},
	"NT": {
		Code:      "NT",
		Name:      "Natural Tiger",
		Price:     17,
		PriceUnit: "cents",
	},
	"DE": {
		Code:      "DE",
		Name:      "Dazzling Elderberry",
		Price:     21,
		PriceUnit: "cents",
	},
	"YR": {
		Code:      "YR",
		Name:      "Yellow Rambutan",
		Price:     23,
		PriceUnit: "cents",
	},
}
