package main

import (
	"fmt"

	"github.com/paulojosecb/go-algorithms/search"
)

func main() {
	sellerNetwork := map[string][]string{
		"you":    []string{"alice", "bob", "claire"},
		"bob":    []string{"anuj", "peggy"},
		"alice":  []string{"peggy"},
		"claire": []string{"thom", "jonny"},
		"anuj":   []string{},
		"peggy":  []string{},
		"thom":   []string{},
		"jonny":  []string{},
	}

	fmt.Println(search.BreadthFirstSearch(sellerNetwork, "you"))
}
