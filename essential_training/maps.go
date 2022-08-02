package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
		Me llamo Juan
		Tengo 27 años
		Soy de Córdoba
		Me gusta Córdoba
		Por que hay mucha
		gente llamada Juan
	`
	words := strings.Fields(text)
	counts := map[string]int{}
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)

}
