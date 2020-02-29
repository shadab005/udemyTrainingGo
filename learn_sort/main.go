package main

import (
	"fmt"
	"sort"
	"strings"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return strings.Compare(p[i], p[j]) < 0
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	group := people{"shadab", "ankit", "kaushik", "naaree"}
	fmt.Println(group)
	//sort.Sort(group)
	fmt.Println(group)

	//sort.StringSlice(group).Sort()
	//fmt.Println(group)

	sort.Strings(group)
	fmt.Println(group)

	sort.Sort(sort.Reverse(sort.StringSlice(group)))
	fmt.Println(group)

	// Simple string sort
	list := []string {"zzz", "asdf", "sdc", "sve"}
	sort.Strings(list)
	fmt.Println(list)

}

