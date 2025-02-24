package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Yusuf Supriadi", "Yusuf"))
	fmt.Println(strings.Split("Yusuf Supriadi", " "))
	fmt.Println(strings.ToLower("Yusuf Supriadi"))
	fmt.Println(strings.ToUpper("Yusuf Supriadi"))
	fmt.Println(strings.Trim("      Yusuf Supriadi     ", " "))
	fmt.Println(strings.ReplaceAll("Yusuf Supriadi Yusuf Supriadi", "Supriadi", "Luchiana"))
}
