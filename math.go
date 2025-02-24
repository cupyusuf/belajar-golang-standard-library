package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.50))
	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Abs(-4))
	fmt.Println(math.Mod(5, 2))
}
