package main

import (
	"day08/ex00"
	"day08/ex01"
	"day08/ex02"
	"log"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	res, _ := ex00.GetElement(arr, 3)
	log.Println("THIS VALUE IS UNSAFELY DEREFERENCED: ", res)

	p1 := ex01.UnknownPlant{
		FlowerType: "rose",
		LeafType:   "cirlce",
		Color:      6969,
	}
	p2 := ex01.AnotherUnknownPlant{
		FlowerColor: 69,
		LeafType:    "Some sorta leaf idk",
		Height:      100,
	}

	ex01.DescribePlants(p1)
	ex01.DescribePlants(p2)

	ex02.CreateWindow("School 21 (by malton)")
}
