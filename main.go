package main

import (
	"fmt"
	"sort"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Slice(fruits, func(i, j int) bool {
		return len(fruits[i]) < len(fruits[j])
	})
	fmt.Println(fruits)
}