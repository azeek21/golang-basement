package main

import (
	"day09/ex00"
	"day09/ex01"
	"day09/ex02"
	"fmt"
)

func main() {
	fmt.Println("EX00: ")
	ex00.Showcase()

	fmt.Println("EX01: ")
	ex01.CrawlWebShowcase([]string{"https://google.com", "https://askaraliev.uz", "https://askaraliev.uz", "https://askaraliev.uz", "https://askaraliev.uz"})

	fmt.Println("EX02: ")
	ex02.ShowCaseMultiplex()
}
