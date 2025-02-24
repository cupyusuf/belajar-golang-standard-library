package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"yusuf", "supriadi"})
	_ = writer.Write([]string{"firlana", "luchiana", "dewi"})
	_ = writer.Write([]string{"shalma", "shabila", "khaerani"})

	writer.Flush()
}
