package examples

import (
	"fmt"
	breadthfirstsearch "grogos/breadth-first-search"
)

type Person struct {
	name        string
	sellsMelons bool
}

func isSeller(p Person) bool {
	return p.sellsMelons
}

type TFriendsGraph = map[string][]Person

func FindMelonSellerBFS() {
	friendsGraph := make(TFriendsGraph)
	friendsGraph["you"] = []Person{
		{name: "Bob", sellsMelons: false},
		{name: "Claire", sellsMelons: false},
		{name: "Alice", sellsMelons: false},
	}
	friendsGraph["Bob"] = []Person{
		{name: "Anuj", sellsMelons: false},
		{name: "Peggy", sellsMelons: false},
	}
	friendsGraph["Claire"] = []Person{
		{name: "Thom", sellsMelons: false},
		{name: "Jonny", sellsMelons: false},
	}
	friendsGraph["Alice"] = []Person{
		{name: "Peggy", sellsMelons: false},
	}
	friendsGraph["Thom"] = []Person{
		{name: "Not seller", sellsMelons: false},
		{name: "Malton", sellsMelons: true},
	}

	res, err := breadthfirstsearch.BreadthFirstSearch(
		friendsGraph,
		func(g TFriendsGraph) []Person {
			return g["you"]
		},
		func(g TFriendsGraph, p Person) []Person {
			return g[p.name]
		},
		isSeller,
	)
	fmt.Printf("[FindMelonSellerBFS]: Seller is: %v, err: %v\n", res, err)
}
