package main

import (
	"fmt"
	"sort"
)

// PersonNum for list
type PersonNum struct {
	Name string
	number  int
}

// List for sorting
type List []PersonNum

// Len is the number of elements in the collection
func (list List) Len() int {
	return len(list)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (list List) Less(i, j int) bool {
	return list[i].Name < list[j].Name
}

// Swap swaps the elements with indexes i and j
func (list List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func main() {
	people := []PersonNum{
		{"Andrey", 89638914356},
		{"Zark", 89632214376},
		{"John", 89638222390},
		{"Luking", 89634364356},
		{"Lisa", 89632214564},
		{"Mary", 89638222344},
	}
	
	sort.Sort(List(people))
	fmt.Println(people)
}
