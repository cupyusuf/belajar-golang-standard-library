package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Yusuf", "Firlana", "Shalma"}
	values := []int{100, 95, 80}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Yusuf"))
	fmt.Println(slices.Index(names, "Yusuf"))
	fmt.Println(slices.Index(names, "Shalma"))
}
