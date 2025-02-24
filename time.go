package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2022, time.December, 31, 23, 59, 59, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	fmt.Println(parse.Local())
}
