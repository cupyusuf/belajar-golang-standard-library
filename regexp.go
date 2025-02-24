package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`yu([a-z])uf`)

	fmt.Println(regex.MatchString("yusuf"))
	fmt.Println(regex.MatchString("yuzuf"))
	fmt.Println(regex.MatchString("yuSuf"))

	fmt.Println(regex.FindAllString("yusuf yosef yisif yasaf yuSuf", 10))
}
