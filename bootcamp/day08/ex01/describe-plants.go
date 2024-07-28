package ex01

import "fmt"

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches" json:"height" binding:"required"`
}

func DescribePlants(plant interface{}) {
	fmt.Print(DescribeStruct(plant))
}
